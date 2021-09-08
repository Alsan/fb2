package cmd

import (
	"crypto/tls"
	"encoding/hex"
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	"github.com/alsan/filebrowser/server/auth"
	"github.com/alsan/filebrowser/server/diskcache"
	"github.com/alsan/filebrowser/server/frontend"
	h "github.com/alsan/filebrowser/server/helpers"
	fbhttp "github.com/alsan/filebrowser/server/http"
	"github.com/alsan/filebrowser/server/img"
	"github.com/alsan/filebrowser/server/rpc"
	"github.com/alsan/filebrowser/server/settings"
	"github.com/alsan/filebrowser/server/storage"
	"github.com/alsan/filebrowser/server/users"
	"github.com/mitchellh/go-homedir"
	"github.com/soheilhy/cmux"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	v "github.com/spf13/viper"
	"google.golang.org/grpc"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)
	cobra.MousetrapHelpText = ""

	rootCmd.SetVersionTemplate("File Browser version {{printf \"%s\" .Version}}\n")

	flags := rootCmd.Flags()
	persistent := rootCmd.PersistentFlags()

	persistent.StringVarP(&cfgFile, "config", "c", "", "config file path")
	persistent.StringP("database", "d", "./filebrowser.db", "database path")
	flags.Bool("noauth", false, "use the noauth auther when using quick setup")
	flags.String("username", "admin", "username for the first user when using quick config")
	flags.String("password", "", "hashed password for the first user when using quick config (default \"admin\")")

	addServerFlags(flags)
}

func addServerFlags(flags *pflag.FlagSet) {
	flags.StringP("address", "a", "127.0.0.1", "address to listen on")
	flags.StringP("log", "l", "stdout", "log output")
	flags.StringP("port", "p", "8080", "port to listen on")
	flags.StringP("cert", "t", "", "tls certificate")
	flags.StringP("key", "k", "", "tls key")
	flags.StringP("root", "r", ".", "root to prepend to relative paths")
	flags.String("socket", "", "socket to listen to (cannot be used with address, port, cert nor key flags)")
	flags.Uint32("socket-perm", 0666, "unix socket file permissions") //nolint:gomnd
	flags.StringP("baseurl", "b", "", "base url")
	flags.String("cache-dir", "", "file cache directory (disabled if empty)")
	flags.Int("img-processors", 4, "image processors count") //nolint:gomnd
	flags.Bool("disable-thumbnails", false, "disable image thumbnails")
	flags.Bool("disable-preview-resize", false, "disable resize of image previews")
	flags.Bool("disable-exec", false, "disables Command Runner feature")
	flags.Bool("disable-type-detection-by-header", false, "disables type detection by reading file headers")
}

func grpcServer(listener net.Listener) {
	s := grpc.NewServer()
	fb.RegisterFileBrowserRpcServiceServer(s, &rpc.Server{})

	err := s.Serve(listener)
	c.ExitIfError("Unable to start grpc server: %v", err)
}

func setupDb(cmd *cobra.Command, d h.PythonData) {
	if !d.HadDB {
		quickSetup(cmd.Flags(), d)
	}

	rpc.SetStorage(&d)
}

func httpServer(listener net.Listener, handler http.Handler) {
	err := http.Serve(listener, handler)
	c.ExitIfError("Unable to start http server: %v", err)
}

func getImgSvc(cmd *cobra.Command) *img.Service {
	workersCount, err := cmd.Flags().GetInt("img-processors")
	c.CheckErr(err)
	if workersCount < 1 {
		log.Fatal("Image resize workers count could not be < 1")
	}

	return img.New(workersCount)
}

func getFileCache(cmd *cobra.Command) diskcache.Interface {
	var fileCache diskcache.Interface = diskcache.NewNoOp()
	cacheDir, err := cmd.Flags().GetString("cache-dir")
	c.CheckErr(err)

	if cacheDir != "" {
		if err := os.MkdirAll(cacheDir, 0700); err != nil {
			log.Fatalf("can't make directory %s: %v", cacheDir, err)
		}
		fileCache = diskcache.New(afero.NewOsFs(), cacheDir)
	}

	return fileCache
}

func getAssetsFs() fs.FS {
	assetsFs, err := fs.Sub(frontend.Assets(), "dist")
	c.ExitIfError("Unable to get assets directory", err)

	return assetsFs
}

func setupHttpHandler(cmd *cobra.Command, d h.PythonData, serverConf *settings.Server) http.Handler {
	imgSvc := getImgSvc(cmd)
	fileCache := getFileCache(cmd)
	assetsFs := getAssetsFs()

	handler, err := fbhttp.NewHandler(imgSvc, fileCache, d.Store, serverConf, assetsFs)
	c.CheckErr(err)

	return handler
}

func captureOsSignalNotification(listener net.Listener) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go cleanupHandler(listener, sigc)
}

func getSocketLister(cmd *cobra.Command, serverConf *settings.Server) net.Listener {
	listener, err := net.Listen("unix", serverConf.Socket)
	c.CheckErr(err)

	socketPerm, err := cmd.Flags().GetUint32("socket-perm") //nolint:govet
	c.CheckErr(err)

	err = os.Chmod(serverConf.Socket, os.FileMode(socketPerm))
	c.CheckErr(err)

	return listener
}

func getTlsListener(serverConf *settings.Server, addr string) net.Listener {
	cer, err := tls.LoadX509KeyPair(serverConf.TLSCert, serverConf.TLSKey) //nolint:govet
	c.CheckErr(err)

	listener, err := tls.Listen("tcp", addr, &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{cer}},
	)
	c.CheckErr(err)

	return listener
}

func getHttpListener(addr string) net.Listener {
	listener, err := net.Listen("tcp", addr)
	c.CheckErr(err)

	return listener
}

func getListener(cmd *cobra.Command, serverConf *settings.Server, addr string) net.Listener {
	if serverConf.Socket != "" {
		return getSocketLister(cmd, serverConf)
	}

	if serverConf.TLSKey != "" && serverConf.TLSCert != "" {
		return getTlsListener(serverConf, addr)
	}

	return getHttpListener(addr)
}

func setupServer(cmd *cobra.Command, d h.PythonData, serverConf *settings.Server, addr string) {
	listener := getListener(cmd, serverConf, addr)

	captureOsSignalNotification(listener)
	defer listener.Close()

	tcpm := cmux.New(listener)
	grpcFilter := tcpm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpFilter := tcpm.Match(cmux.Any())
	handler := setupHttpHandler(cmd, d, serverConf)

	go grpcServer(grpcFilter)
	go httpServer(httpFilter, handler)

	log.Printf("server listenering on port :%s", addr)
	if err := tcpm.Serve(); err != nil {
		log.Fatalf("Error serving cmux: %v", err)
	}
}

func setupServerConf(conf *settings.Server) {
	setupLog(conf.Log)

	root, err := filepath.Abs(conf.Root)
	c.CheckErr(err)
	conf.Root = root
}

var rootCmd = &cobra.Command{
	Use:   "filebrowser",
	Short: "A stylish web-based file browser",
	Long: `File Browser CLI lets you create the database to use with File Browser,
manage your users and all the configurations without acessing the
web interface.

If you've never run File Browser, you'll need to have a database for
it. Don't worry: you don't need to setup a separate database server.
We're using Bolt DB which is a single file database and all managed
by ourselves.

For this specific command, all the flags you have available (except
"config" for the configuration file), can be given either through
environment variables or configuration files.

If you don't set "config", it will look for a configuration file called
.filebrowser.{json, toml, yaml, yml} in the following directories:

- ./
- $HOME/
- /etc/filebrowser/

The precedence of the configuration values are as follows:

- flags
- environment variables
- configuration file
- database values
- defaults

The environment variables are prefixed by "FB_" followed by the option
name in caps. So to set "database" via an env variable, you should
set FB_DATABASE.

Also, if the database path doesn't exist, File Browser will enter into
the quick setup mode and a new database will be bootstraped and a new
user created with the credentials from options "username" and "password".`,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		serverConf := getRunParams(cmd.Flags(), d.Store)
		addr := serverConf.Address + ":" + serverConf.Port

		setupServerConf(serverConf)
		setupDb(cmd, d)
		setupServer(cmd, d, serverConf, addr)
	}, h.PythonConfig{AllowNoDB: true}),
}

func cleanupHandler(listener net.Listener, c chan os.Signal) { //nolint:interfacer
	sig := <-c
	log.Printf("Caught signal %s: shutting down.", sig)
	listener.Close()
	os.Exit(0)
}

//nolint:gocyclo
func getRunParams(flags *pflag.FlagSet, st *storage.Storage) *settings.Server {
	server, err := st.Settings.GetServer()
	c.CheckErr(err)

	if val, set := h.GetParamB(flags, "root"); set {
		server.Root = val
	}

	if val, set := h.GetParamB(flags, "baseurl"); set {
		server.BaseURL = val
	}

	if val, set := h.GetParamB(flags, "log"); set {
		server.Log = val
	}

	isSocketSet := false
	isAddrSet := false

	if val, set := h.GetParamB(flags, "address"); set {
		server.Address = val
		isAddrSet = isAddrSet || set
	}

	if val, set := h.GetParamB(flags, "port"); set {
		server.Port = val
		isAddrSet = isAddrSet || set
	}

	if val, set := h.GetParamB(flags, "key"); set {
		server.TLSKey = val
		isAddrSet = isAddrSet || set
	}

	if val, set := h.GetParamB(flags, "cert"); set {
		server.TLSCert = val
		isAddrSet = isAddrSet || set
	}

	if val, set := h.GetParamB(flags, "socket"); set {
		server.Socket = val
		isSocketSet = isSocketSet || set
	}

	if isAddrSet && isSocketSet {
		c.CheckErr(errors.New("--socket flag cannot be used with --address, --port, --key nor --cert"))
	}

	// Do not use saved Socket if address was manually set.
	if isAddrSet && server.Socket != "" {
		server.Socket = ""
	}

	_, disableThumbnails := h.GetParamB(flags, "disable-thumbnails")
	server.EnableThumbnails = !disableThumbnails

	_, disablePreviewResize := h.GetParamB(flags, "disable-preview-resize")
	server.ResizePreview = !disablePreviewResize

	_, disableTypeDetectionByHeader := h.GetParamB(flags, "disable-type-detection-by-header")
	server.TypeDetectionByHeader = !disableTypeDetectionByHeader

	_, disableExec := h.GetParamB(flags, "disable-exec")
	server.EnableExec = !disableExec

	return server
}

func setupLog(logMethod string) {
	switch logMethod {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	case "":
		log.SetOutput(ioutil.Discard)
	default:
		log.SetOutput(&lumberjack.Logger{
			Filename:   logMethod,
			MaxSize:    100,
			MaxAge:     14,
			MaxBackups: 10,
		})
	}
}

func quickSetup(flags *pflag.FlagSet, d h.PythonData) {
	set := &settings.Settings{
		Key:           generateKey(),
		Signup:        false,
		CreateUserDir: false,
		Defaults: settings.UserDefaults{
			Scope:       ".",
			Locale:      "en",
			SingleClick: false,
			Perm: users.Permissions{
				Admin:    false,
				Execute:  true,
				Create:   true,
				Rename:   true,
				Modify:   true,
				Delete:   true,
				Share:    true,
				Download: true,
			},
		},
	}

	var err error
	if _, noauth := h.GetParamB(flags, "noauth"); noauth {
		set.AuthMethod = auth.MethodNoAuth
		err = d.Store.Auth.Save(&auth.NoAuth{})
	} else {
		set.AuthMethod = auth.MethodJSONAuth
		err = d.Store.Auth.Save(&auth.JSONAuth{})
	}

	c.CheckErr(err)
	err = d.Store.Settings.Save(set)
	c.CheckErr(err)

	ser := &settings.Server{
		BaseURL: h.GetParam(flags, "baseurl"),
		Port:    h.GetParam(flags, "port"),
		Log:     h.GetParam(flags, "log"),
		TLSKey:  h.GetParam(flags, "key"),
		TLSCert: h.GetParam(flags, "cert"),
		Address: h.GetParam(flags, "address"),
		Root:    h.GetParam(flags, "root"),
	}

	err = d.Store.Settings.SaveServer(ser)
	c.CheckErr(err)

	username := h.GetParam(flags, "username")
	password := h.GetParam(flags, "password")

	if password == "" {
		password = hex.EncodeToString(c.Md5Pass("admin"))
	}

	if username == "" || password == "" {
		log.Fatal("username and password cannot be empty during quick setup")
	}

	user := &users.User{
		Username:     username,
		Password:     password,
		LockPassword: false,
	}

	set.Defaults.Apply(user)
	user.Perm.Admin = true

	err = d.Store.Users.Save(user)
	c.CheckErr(err)
}

func initConfig() {
	if cfgFile == "" {
		home, err := homedir.Dir()
		c.CheckErr(err)
		v.AddConfigPath(".")
		v.AddConfigPath(home)
		v.AddConfigPath("/etc/filebrowser/")
		v.SetConfigName("filebrowser")
	} else {
		v.SetConfigFile(cfgFile)
	}

	v.SetEnvPrefix("FB")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)

		if _, ok := err.(v.ConfigParseError); ok {
			panic(err)
		}
		cfgFile = "No config file used"
	} else {
		cfgFile = "Using config file: " + v.ConfigFileUsed()
	}
}

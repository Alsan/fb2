package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// AuthClient is a client to call authentication RPC
type authClient struct {
	service  fb.FileBrowserRpcServiceClient
	username string
	password string
}

// NewAuthClient returns a new auth client
func newAuthClient(conn *grpc.ClientConn, username, password string) *authClient {
	service := fb.NewFileBrowserRpcServiceClient(conn)
	return &authClient{service, username, password}
}

// Login login user and returns the access token
func (client *authClient) login() (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &fb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return false, "", err
	}

	if res.GetStatus() == fb.ReplyStatus_Ok {
		return true, res.GetToken(), nil
	}

	return false, res.GetMessage(), nil
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func getUserServerInput() string {
	server := c.GetUserInput("Server (localhost:8080)")

	if server == "" {
		server = "localhost:8080"
	}

	return server
}

func encryptPassword(password string) string {
	m := c.Md5Pass(password)
	h := hex.EncodeToString(m)
	b := []byte(h)

	return string(c.BcryptHash(b))
}

func doLogin(server, username, password string) (string, bool) {
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
	c.ExitIfError("Unable to connect to server, %v", err)
	defer conn.Close()

	client := newAuthClient(conn, username, password)
	success, msg, err := client.login()
	c.CheckErr(err)

	return msg, success
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into server",
	Long:  `Login into filebrowser server`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		server := getUserServerInput()
		username := c.GetUserInput("Username")
		password := encryptPassword(c.GetUserPasswordInput())

		msg, ok := doLogin(server, username, password)

		if !ok {
			log.Fatalf("Login failed: %s\n", msg)
		}

		fmt.Printf("Token: %s\n", msg)
	},
}

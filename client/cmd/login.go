package cmd

import (
	"context"
	"log"
	"time"

	"github.com/alsan/filebrowser/client/utils"
	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

func getServer() string {
	server, err := utils.GetUserInput("Server (localhost:8080)")
	c.CheckErr(err)

	if server == "" {
		server = "localhost:8080"
	}

	return server
}

func getUsername() string {
	username, err := utils.GetUserInput("Username")
	c.CheckErr(err)

	return username
}

func getPassword() string {
	password, err := utils.GetUserPasswordInput()
	c.CheckErr(err)

	return string(c.BcryptHash(c.Md5Pass(password)))
}

func doLogin(server string, username string, password string) *fb.LoginReply {
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
	c.ExitIfError("Unable to connect to server, %v", err)
	defer conn.Close()
	client := fb.NewFileBrowserRpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := client.Login(ctx, &fb.LoginRequest{
		Username: username,
		Password: password,
	})
	c.ExitIfError("Unable to login, %v", err)

	return reply
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into server",
	Long:  `Login into filebrowser server`,
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		server := getServer()
		username := getUsername()
		password := getPassword()

		reply := doLogin(server, username, password)

		if reply.GetStatus() == fb.ReplyStatus_Ok {
			log.Printf("Login successful, Token: %s", reply.GetToken())
		} else {
			log.Printf("Login failed, message: %s", reply.GetMessage())
		}
	},
}

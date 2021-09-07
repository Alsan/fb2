package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	clientUtils "github.com/alsan/filebrowser/client/utils"
	fb "github.com/alsan/filebrowser/proto"
	utils "github.com/alsan/filebrowser/utils"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

func getServer() string {
	server, err := clientUtils.GetUserInput("Server (localhost:8080)")
	utils.CheckErr(err)

	if server == "" {
		server = "localhost:8080"
	}

	return server
}

func getUsername() string {
	username, err := clientUtils.GetUserInput("Username")
	utils.CheckErr(err)

	return username
}

func getPassword() string {
	password, err := clientUtils.GetUserPasswordInput()
	utils.CheckErr(err)

	return password
}

func doLogin(server string, username string, password string) *fb.LoginReply {
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
	utils.ExitIfError("Unable to connect to server, %v", err)
	defer conn.Close()
	client := fb.NewFileBrowserRpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := client.Login(ctx, &fb.LoginRequest{
		Username: username,
		Password: password,
	})
	utils.ExitIfError("Unable to login, %v", err)

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
		md5Password := utils.Md5Pass(password)

		fmt.Printf("md5Password: %s\n", md5Password)

		reply := doLogin(server, username, password)

		if reply.GetStatus() == fb.ReplyStatus_Ok {
			log.Printf("Login successful, Token: %s", reply.GetToken())
		} else {
			log.Printf("Login failed, message: %s", reply.GetMessage())
		}
	},
}

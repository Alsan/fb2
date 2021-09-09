package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	u "github.com/alsan/filebrowser/client/utils"
	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

func getServer() string {
	server, err := u.GetUserInput("Server (localhost:8080)")
	c.CheckErr(err)

	if server == "" {
		server = "localhost:8080"
	}

	return server
}

func getUsername() string {
	username, err := u.GetUserInput("Username")
	c.CheckErr(err)

	return username
}

func getPassword() string {
	password, err := u.GetUserPasswordInput()
	c.CheckErr(err)

	m := c.Md5Pass(password)
	h := hex.EncodeToString(m)
	b := []byte(h)

	return string(c.BcryptHash(b))
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
	c.CheckErr(err)

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
			fmt.Printf("Login successful, Token: %s", reply.GetToken())
		} else {
			fmt.Printf("Login failed, message: %s", reply.GetMessage())
		}

		fmt.Println()
	},
}

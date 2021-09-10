package cmd

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/alsan/filebrowser/client/rpc"
	c "github.com/alsan/filebrowser/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(loginCmd)
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

	client := rpc.NewAuthClient(conn, username, password)
	return client.Login()
}

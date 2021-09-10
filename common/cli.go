package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/pflag"
	"golang.org/x/term"
)

func MustGetString(flags *pflag.FlagSet, flag string) string {
	s, err := flags.GetString(flag)
	CheckErr(err)
	return s
}

func MustGetBool(flags *pflag.FlagSet, flag string) bool {
	b, err := flags.GetBool(flag)
	CheckErr(err)
	return b
}

func MustGetUint(flags *pflag.FlagSet, flag string) uint {
	b, err := flags.GetUint(flag)
	CheckErr(err)
	return b
}

func GetUserInput(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", label)
	text, err := reader.ReadString('\n')
	CheckErr(err)

	return strings.Replace(text, "\n", "", -1)
}

func GetUserPasswordInput() string {
	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	CheckErr(err)

	fmt.Println("")
	password := string(bytePassword)
	return strings.TrimSpace(password)
}

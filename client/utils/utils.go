package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func GetUserInput(label string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", label)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Replace(text, "\n", "", -1), nil
}

func GetUserPasswordInput() (string, error) {
	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	fmt.Println("")
	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}

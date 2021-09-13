package common

import (
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExitIfError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

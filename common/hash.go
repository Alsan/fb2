package common

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Reverse(str string) string {
	n := len(str)
	runes := make([]rune, n)

	for _, rune := range str {
		n--
		runes[n] = rune
	}

	return string(runes[n:])
}

func Md5sum(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}

// Md5Pass calculate password hash
// 1. reverse the password, ie: admin -> nimda
// 2. split the reverse password into array
// 3. join the array with comma as a string, ie: n,i,m,d,a
// 4. calculate md5 checksum of the result string
func Md5Pass(password string) string {
	s := strings.Join(strings.Split(Reverse(password), ""), ",")

	return Md5sum(s)
}
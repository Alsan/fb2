package common

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Reverse(str string) []byte {
	n := len(str)
	runes := make([]rune, n)

	for _, rune := range str {
		n--
		runes[n] = rune
	}

	return []byte(string(runes[n:]))
}

func Md5sum(password []byte) []byte {
	hash := md5.Sum(password)

	return hash[:]
}

// Md5Pass calculate password hash
// 1. reverse the password, ie: admin -> nimda
// 2. split the reverse password into array
// 3. join the array with comma as a string, ie: n,i,m,d,a
// 4. calculate md5 checksum of the result string
func Md5Pass(password string) []byte {
	r := string(Reverse(password))
	s := strings.Join(strings.Split(r, ""), ",")
	b := Md5sum([]byte(s))

	return b
}

func BcryptHash(password []byte) []byte {
	bytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	CheckErr(err)

	return bytes
}

func GetFileChecksum(file *os.File) string {
	// reset file pointer to begining of the file
	file.Seek(0, io.SeekStart)

	hasher := md5.New()
	_, err := io.Copy(hasher, file)
	CheckErr(err)

	return hex.EncodeToString(hasher.Sum(nil))
}

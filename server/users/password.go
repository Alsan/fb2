package users

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func reverse(str string) string {
	n := len(str)
	runes := make([]rune, n)

	for _, rune := range str {
		n--
		runes[n] = rune
	}

	return string(runes[n:])
}

func md5sum(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}

// Md5Pass calculate password hash
// 1. reverse the password, ie: admin -> nimda
// 2. split the reverse password into array
// 3. join the array with comma as a string, ie: n,i,m,d,a
// 4. calculate md5 checksum of the result string
func Md5Pass(password string) string {
	s := strings.Join(strings.Split(reverse(password), ""), ",")

	return md5sum(s)
}

// HashPwd hashes a password.
func HashPwd(password string) (string, error) {
  // we don't really hash the password
  return password, nil
}

// CheckPwd checks if a password is correct.
func CheckPwd(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))

	if err != nil {
		log.Error().Err(err).Msg("bcrypt compare hash and password failed")
	}

	return err == nil
}

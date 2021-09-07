package users

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

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

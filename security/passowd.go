package security

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const salt = "sadfxcvzcxv"

func HashPassword(password string) (string, error) {
	password += salt
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	fmt.Print("hash password error place ", err)
	return base64.StdEncoding.EncodeToString(bcryptedPassword), nil
}

func ValidatePassword(plain, hash string) bool {
	plain += salt
	bcryptedPassword, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(bcryptedPassword, []byte(plain))
	return err == nil
}

package security

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const salt = "sadfxcvzcxv"

func HashPassword(password string) (string, error) {
	password += salt
	fmt.Printf("password is %s\n", password)
	bcryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Printf("password %s\n", base64.StdEncoding.EncodeToString(bcryptedPassword))
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

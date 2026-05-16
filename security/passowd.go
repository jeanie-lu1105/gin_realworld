package security

import (
	"encoding/base64"

	"example.com/gin_realworld/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const salt = "sadfxcvzcxv"

func HashPassword(password string) (string, error) {
	password += salt
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString([]byte(bcryptedPassword)), nil
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

func VerifyJWT(token string) (*jwt.MapClaims, bool, error) {
	var claim jwt.MapClaims
	claims, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetSecret()), nil
	})
	if err != nil {
		return nil, false, err
	}
	if claims.Valid {
		return &claim, true, nil
	}
	return nil, false, nil
}

func GetCurrentUsername(ctx *gin.Context) string {
	mapClaims := ctx.MustGet("user").(*jwt.MapClaims)
	username := (*mapClaims)["user"].(map[string]interface{})["username"].(string)
	return username
}

func GetCurrentUserEmail(ctx *gin.Context) string {
	mapClaims := ctx.MustGet("user").(*jwt.MapClaims)
	email := (*mapClaims)["user"].(map[string]interface{})["email"].(string)
	return email
}

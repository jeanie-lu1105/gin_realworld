package security

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"example.com/gin_realworld/config"
	"example.com/gin_realworld/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	var err error
	var bytes []byte
	bytes, err = os.ReadFile(config.GetPrivateKeyLocation())
	if err != nil {
		panic("private key not found")
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}

	bytes, err = os.ReadFile(config.GetPublicKeyLocation())
	if err != nil {
		panic("public key not found")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}

}

func GenerateJWT(username, email string) (string, error) {
	key := []byte(config.GetSecret())
	tokenDuration := 24 * time.Hour
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": map[string]string{
				"email":    email,
				"username": username,
			},
			"iat": now.Unix(),
			"exp": now.Add(tokenDuration).Unix(),
		})

	return t.SignedString(key)
}

func VerifyJWT(token string) (*jwt.MapClaims, bool, error) {
	var claim jwt.MapClaims
	claims, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetSecret()), nil
	})
	fmt.Printf("verify jwt claims %s\n", utils.JsonMarshal(claims))
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

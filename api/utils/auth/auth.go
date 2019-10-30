package auth

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJSONWebToken(id int) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = id
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Create signed token
	tokenString, err := token.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		log.Println(err.Error())
	}

	return tokenString, err
}

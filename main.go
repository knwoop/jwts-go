package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var mySigningKey = []byte(os.Getenv("MY_JWT_TOKEN"))
//var mySigningKey = []byte("mysupersecret")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true;
	claims["user"]       = "Taro Forbes"
	claims["exp"]        = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	fmt.Println("My simple client")
	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token string")
	}

	fmt.Println(tokenString)
}

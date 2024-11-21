package handler

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func createToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"password": password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getClaimsFromToken(tokenString string) (string, string, error) {
	var username string
	var password string
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username = fmt.Sprint(claims["username"])
		password = fmt.Sprint(claims["password"])
	}

	if username == "" || password == "" {
		return "", "", fmt.Errorf("invalid token payload")
	}

	return username, password, nil
}

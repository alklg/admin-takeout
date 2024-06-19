package login

import (
	"crypto/rand"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func GenerateJwt(key any, method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString(key)
}

func NewJwt() (string, error) {
	jwtKey := make([]byte, 32)

	_, err := rand.Read(jwtKey)

	if err != nil {
		log.Printf("rand byte slice error = %v\n", err)
		return "", err
	}

	jwtStr, err := GenerateJwt(jwtKey, jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "nidielkl",
		"sub": "admin-takeout",
		"aud": "employee",
	})

	if err != nil {
		log.Printf("new jwt error = %v\n", err)
		return "", err
	}

	log.Printf("this is jwtStr %v\n", jwtStr)
	return jwtStr, nil
}

package token

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(claim *Claims, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	strToken, err := token.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}
	return strToken, nil
}

func ParseJWT(strToken string, key []byte) (*Claims, error) {
	claim := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(strToken, claim, func(strToken *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if !parsedToken.Valid {
		log.Fatal("Токен говно") // ну давай заплачь
	}
	return claim, nil
}

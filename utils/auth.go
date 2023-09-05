package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken() (string, error) {
	secretKey := GetConfig("JWT_SECRET_KEY")
	jwtExpiry, _ := strconv.Atoi(GetConfig("JWT_SECRET_KEY_EXPIRY"))

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(jwtExpiry)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

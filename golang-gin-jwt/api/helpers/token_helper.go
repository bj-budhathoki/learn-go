package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int, name, email, role, secret string, expiryIn time.Duration) (string, error) {
	tokenByte := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["exp"] = now.Add(expiryIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["name"] = name
	claims["role"] = role
	claims["email"] = email
	tokenString, err := tokenByte.SignedString([]byte(secret))
	return tokenString, err
}

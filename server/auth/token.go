package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(secret string, claims map[string]any, expiryHours int) (string, error) {
	stdClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(expiryHours) * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	for k, v := range claims {
		stdClaims[k] = v
	}

	// Create token with HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)

	return token.SignedString([]byte(secret))
}

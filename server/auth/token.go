package auth

import (
	"maps"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(secret string, claims map[string]any, expiryHours int) (string, error) {
	stdClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(expiryHours) * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	maps.Copy(stdClaims, claims)

	// Create token with HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)

	return token.SignedString([]byte(secret))
}

func ParseToken(secret string, tokenStr string) (map[string]any, error) {
	parser := jwt.NewParser(jwt.WithLeeway(1 * time.Minute))

	token, err := parser.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// get claims map
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

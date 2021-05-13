package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func Generate(claims jwt.MapClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func Verify(token string, secretKey string) (*jwt.Token, jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, nil, err
	}

	return t, t.Claims.(jwt.MapClaims), nil
}

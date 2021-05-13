package token

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func generateJWT(claims jwtgo.MapClaims, exp time.Duration, secret string) (string, error) {
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = exp

	return jwt.Generate(claims, secret)
}

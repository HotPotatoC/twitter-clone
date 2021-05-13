package token

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/config"
	"github.com/HotPotatoC/twitter-clone/internal/common/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
)

type AccessToken struct {
	token string
	expAt time.Time
}

func NewAccessToken(claims jwtgo.MapClaims) (*AccessToken, error) {
	exp := config.GetDuration("ACCESS_TOKEN_DURATION",
		time.Duration(time.Now().Add(time.Minute*15).Unix()))
	secret := config.GetString("ACCESS_TOKEN_SECRET", "")

	token, err := generateJWT(claims, exp, secret)
	if err != nil {
		return nil, err
	}

	at := new(AccessToken)
	at.token = token
	at.expAt = time.Now().Add(config.GetDuration("ACCESS_TOKEN_DURATION",
		time.Duration(time.Now().Add(time.Minute*15).Unix())))

	return at, nil
}

func VerifyAccessToken(tokenStr string) (jwtgo.MapClaims, error) {
	token, claims, err := jwt.Verify(tokenStr, config.GetString("ACCESS_TOKEN_SECRET", ""))
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwtgo.Claims); !ok && !token.Valid {
		return nil, jwtgo.ErrInvalidKey
	}

	return claims, nil
}

func (t *AccessToken) String() string {
	return t.token
}

func (t *AccessToken) ExpiresAt() time.Time {
	return t.expAt
}

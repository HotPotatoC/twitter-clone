package token

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/config"
	"github.com/HotPotatoC/twitter-clone/internal/common/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
)

type RefreshToken struct {
	token string
	expAt time.Time
}

func NewRefreshToken(claims jwtgo.MapClaims) (*RefreshToken, error) {
	exp := config.GetDuration("REFRESH_TOKEN_DURATION",
		time.Duration(time.Now().Add(time.Hour*24*7).Unix()))
	secret := config.GetString("REFRESH_TOKEN_SECRET", "")

	token, err := generateJWT(claims, exp, secret)
	if err != nil {
		return nil, err
	}

	rt := new(RefreshToken)
	rt.token = token
	rt.expAt = time.Now().Add(config.GetDuration("REFRESH_TOKEN_DURATION",
		time.Duration(time.Now().Add(time.Hour*24*7).Unix())))

	return rt, nil
}

func VerifyRefreshToken(tokenStr string) (jwtgo.MapClaims, error) {
	token, claims, err := jwt.Verify(tokenStr, config.GetString("REFRESH_TOKEN_SECRET", ""))
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwtgo.Claims); !ok && !token.Valid {
		return nil, jwtgo.ErrInvalidKey
	}

	return claims, nil
}

func (t *RefreshToken) String() string {
	return t.token
}

func (t *RefreshToken) ExpiresAt() time.Time {
	return t.expAt
}

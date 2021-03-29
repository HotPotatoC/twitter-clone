package token

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/pkg/config"
	"github.com/HotPotatoC/twitter-clone/pkg/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type AccessToken struct {
	token string
	expAt time.Time
}

func NewAccessToken(claims jwtgo.MapClaims) (*AccessToken, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = config.GetDuration("ACCESS_TOKEN_DURATION",
		time.Duration(time.Now().Add(time.Minute*15).Unix()))

	token, err := jwt.Generate(claims, config.GetString("ACCESS_TOKEN_SECRET", ""))
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

type RefreshToken struct {
	token string
	expAt time.Time
}

func NewRefreshToken(claims jwtgo.MapClaims) (*RefreshToken, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = config.GetDuration("REFRESH_TOKEN_DURATION",
		time.Duration(time.Now().Add(time.Hour*24*7).Unix()))

	token, err := jwt.Generate(claims, config.GetString("REFRESH_TOKEN_SECRET", ""))
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

package jwt

import (
	"errors"
	"fmt"
	accountmodels "synapsis-test/app/account/domain/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	ErrTokenInvalid = errors.New("token invalid")
)

func EncodeToken(user accountmodels.User, secret string, day int) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * time.Duration(day))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId:   user.Id,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	})

	return token.SignedString([]byte(secret))
}

func DecodeToken(token string, secret string) (Claims, error) {
	claims := &Claims{}
	tokenDecode, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return Claims{}, err
	}
	if !tokenDecode.Valid {
		return Claims{}, ErrTokenInvalid
	}
	return *claims, nil
}

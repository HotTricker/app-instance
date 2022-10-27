package gojwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TokenExpireDuration = time.Minute * 2

var JwtSecret = []byte("HottrickerPassword")

type UserInfo struct {
	Id   int
	Name string
}

type userClaims struct {
	User UserInfo
	jwt.StandardClaims
}

func GenerateToken(userInfo UserInfo) (string, error) {
	expirationTime := time.Now().Add(TokenExpireDuration)
	claims := &userClaims{
		User: userInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "HotTricker",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString(JwtSecret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}

}

func ParseToken(tokenString string) (*userClaims, error) {
	claims := &userClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	return claims, err
}

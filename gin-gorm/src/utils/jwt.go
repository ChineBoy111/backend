package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	Id                   int
	Username             string
	jwt.RegisteredClaims // 隐式声明 RegisteredClaims jwt.RegisteredClaims
}

func GenToken(id int, username string) (token string, err error) {
	expiration := viper.GetDuration("jwt.expiration")

	registeredClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration * time.Second)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   "Token",
	}

	claims := Claims{
		Id:               id,
		Username:         username,
		RegisteredClaims: registeredClaims,
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return unsignedToken.SignedString( /* 强制类型转换 */ []byte(viper.GetString("jwt.signingKey")))
}

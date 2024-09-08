package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

// Payload 载荷
type Payload struct {
	Id                   int
	Username             string
	jwt.RegisteredClaims //! 等价于 RegisteredClaims jwt.RegisteredClaims
}

// []byte 类型的签名密钥
var signingKey = []byte(viper.GetString("jwt.signingKey"))

// GenToken 返回已签名的 tokStr 字符串和可能的错误 err
func GenToken(id int, username string) (tokStr string, err error) {
	expiration := viper.GetDuration("jwt.expiration")

	registeredClaims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),                               //! jwt 令牌签发时间
		Subject:   "Token",                                                      //! jwt 令牌主题
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration * time.Second)), //! jwt 令牌过期时间
	}

	payload := Payload{
		Id:               id,
		Username:         username,
		RegisteredClaims: registeredClaims,
	}

	//! 创建未签名的 unsignedToken，指定签名算法为 HS256，载荷为 payload
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//! 使用签名算法 HS256，[]byte 类型的签名密钥对 unsignedToken 签名
	//! 返回已签名的 tokStr 字符串和可能的错误 err
	return unsignedToken.SignedString( /* 强制类型转换 */ signingKey)
}

// ParseToken 遍历 tokStr 字符串，返回解析出的载荷 payload，可能的错误 err 和 token是否有效 isValid
func ParseToken(tokStr string) (payload Payload, err error, isValid bool) {
	payload = Payload{}                                          // 解析出的载荷
	keyProvider := func(token *jwt.Token) (interface{}, error) { // 提供签名密钥的函数
		return signingKey, nil
	}

	tok, err := jwt.ParseWithClaims(tokStr, &payload, keyProvider)
	if err != nil {
		return payload, err, false
	}
	if !tok.Valid {
		err = errors.New("invalid token")
		return payload, err, false
	}
	return payload, nil, true
}

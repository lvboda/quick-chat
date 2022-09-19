package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lvboda/quick-chat/utils/status"
)

type claims struct {
	Uid      string `json:"userId"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// CreateToken 生成token
func CreateToken(uid string, password string) (string, error) {
	serverConf := GetConfig().Server

	claims := claims{
		Uid:      uid,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + serverConf.TokenAging,
			Issuer:    "quick-chat",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(serverConf.JwtKey))
}

// ParseToken 解析token
func ParseToken(tokenStr string) (any, int) {
	serverConf := GetConfig().Server

	token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(serverConf.JwtKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, status.ERROR_TOKEN_RUNTIME
			} else {
				return nil, status.ERROR_TOKEN_WRONG
			}
		}
	}

	if token == nil {
		return nil, status.ERROR_TOKEN_WRONG
	}

	return token.Claims, status.SUCCESS
}

package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("TodoList")

type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

// GenerateToken 签发用户token
func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todoList",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 首先获取tokenClaims，但是这个还不能拿给前端使用 需要进行加密
	token, err := tokenClaims.SignedString(jwtSecret)                // 加密后获得token
	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

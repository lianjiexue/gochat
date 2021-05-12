package lib

import (
	"app/config"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

func CreateToken(uid int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{uid, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 15000,
		Issuer:    "daguozhensi",
	}})
	tokenString, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		log.Println("生成token失败")
		return ""
	}
	return tokenString
}
func ParseToken(tokenString string) (int, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})
	if err != nil {
		log.Println(err)
		return 0, false
	}
	if claim, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		log.Println(claim.Uid, claim.StandardClaims.ExpiresAt)
		return claim.Uid, true
	} else {
		return 0, false
	}
}

package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var jwtSecret []byte

type Claims struct {
	UID uint
	jwt.StandardClaims
}

func GenerateToken(uid uint) (string, error) {
	jwtSecret = []byte(viper.GetString("app.jwtSecret"))
	now := time.Now()
	expireTime := now.Add(time.Hour * time.Duration(viper.GetInt("app.jwtTimeout")))
	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    viper.GetString("app.name"),
			IssuedAt:  now.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	jwtSecret = []byte(viper.GetString("app.jwtSecret"))
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
func GetUID(token string) (uint, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return 0, err
	}
	return claims.UID, nil
}

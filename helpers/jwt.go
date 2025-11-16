package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint) (string, error) {
	secret := viper.GetString("JWT.Secret")
	issuer := viper.GetString("JWT.Issuer")

	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

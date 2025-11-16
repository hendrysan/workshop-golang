package middlewares

import (
	"net/http"
	"strings"
	"workshop1/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, helpers.Response(http.StatusUnauthorized, "Unauthorized", "Missing or invalid Authorization header", nil))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		secret := viper.GetString("JWT.Secret")

		token, err := jwt.ParseWithClaims(tokenString, &helpers.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, helpers.Response(http.StatusUnauthorized, "Unauthorized", "Invalid token", nil))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*helpers.JWTClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, helpers.Response(http.StatusUnauthorized, "Unauthorized", "Invalid token claims", nil))
			c.Abort()
			return
		}

		// simpan user id di context
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}

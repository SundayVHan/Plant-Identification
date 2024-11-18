package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"plant_identification/config"
	"plant_identification/internal/common"
	"plant_identification/internal/user"
	"plant_identification/internal/util"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println(c.FullPath())
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			common.Abort(c, common.ErrUnauthorized, "jwt not found in header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

		claims := &util.UserClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			common.Abort(c, common.ErrTokenInvalid, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		u, err := user.GetUser(claims.Username)
		if authHeader == "" {
			common.Abort(c, common.ErrUserNotFound, "token valid but user not found", http.StatusInternalServerError)
			return
		}
		c.Set("user", u)

		c.Next()
	}
}

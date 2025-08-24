package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rohan/go-todo/internal/loggers"
)

func JWTAuthMiddlewares() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			loggers.HandleResponse(ctx, http.StatusUnauthorized, "Not Authorized")
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			// Use secret from env
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			loggers.HandleResponse(ctx, http.StatusUnauthorized, "Not Authorized")
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("email", claims["sub"])
			ctx.Set("userId", claims["userId"])
		}

		ctx.Next()

	}
}

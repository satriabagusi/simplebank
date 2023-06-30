/*
Author: Satria Bagus(satria.bagus18@gmail.com)
auth_middleware.go (c) 2023
Desc: description
Created:  2023-06-30T04:22:54.377Z
Modified: !date!
*/

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tokenJwt "github.com/satriabagusi/simplebank/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		payload, err := tokenJwt.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid or expired token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("payload", payload)
		ctx.Next()
	}
}

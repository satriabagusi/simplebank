/*
Author: Satria Bagus(satria.bagus18@gmail.com)
auth_middleware.go (c) 2023
Desc: description
Created:  2023-06-30T04:22:54.377Z
Modified: !date!
*/

package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/simplebank/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		log.Println(tokenString)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		payload, err := token.ValidateToken(tokenString)
		log.Println(err)
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

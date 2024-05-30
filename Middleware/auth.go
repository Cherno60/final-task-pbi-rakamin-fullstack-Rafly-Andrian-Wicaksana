package middleware

import "github.com/gin-gonic/gin"

func AuuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//midleware
	}
}

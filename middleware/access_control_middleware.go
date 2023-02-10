package middleware

import "github.com/gin-gonic/gin"

func AllowAccessMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "localhost:8080")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
}

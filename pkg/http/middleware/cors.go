package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS replies to request with cors header and handles preflight request
// it is enhancement to improve middleware usability instead of wrapping every handler
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT,OPTIONS,HEAD")
		c.Next()
	}
}

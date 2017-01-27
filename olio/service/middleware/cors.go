package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type OlioCORSMiddleware struct {}

func NewOlioCORSMiddleware() OlioCORSMiddleware {
	return OlioCORSMiddleware{}
}

func (m OlioCORSMiddleware) Create() gin.HandlerFunc {
	// from https://github.com/gin-gonic/gin/issues/29#issuecomment-89132826
	return func (c *gin.Context) {
		if c.Request.Header["Origin"] != nil {
			c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header["Origin"][0]) // allow any origin domain
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // allow any origin domain
		}
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://domain.com") // uncomment to restrict to single domain
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Set-Request-Id")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
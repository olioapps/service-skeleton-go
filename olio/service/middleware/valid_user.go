package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidUser(c *gin.Context) {
	_, whitelisted := c.Get("whitelisted")
	if whitelisted {
		c.Next()
		return
	}

	user, exists := c.Get("currentUser")
	if !exists || user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

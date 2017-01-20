package middleware

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type OlioBasicAuthMiddleware struct {
	userExtractor UserExtractor
}

func NewOlioBasicAuthMiddleware(userExtractor UserExtractor) OlioBasicAuthMiddleware {
	middleware := OlioBasicAuthMiddleware{userExtractor}

	return middleware
}

func (m OlioBasicAuthMiddleware) CreateMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// if the user exists, continue
		_, exists := c.Get("whitelisted")
		if exists {
			c.Next()
			return
		}

		{
			user, exists := c.Get("currentUser")
			if exists && user != nil {
				c.Next()
				return
			}
		}

		authHeader := c.Request.Header.Get("Authorization")

		// skip this handler if no authorization header is found
		if authHeader == "" {
			c.Next()
			return
		}

		// otherwise keep processing
		r, _ := regexp.Compile("^Basic (.+)$")

		match := r.FindStringSubmatch(authHeader)

		if len(match) == 0 {
			c.Next()
			return
		}
		tokenString := match[1]

		if len(tokenString) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		str, err := base64.StdEncoding.DecodeString(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		creds := bytes.SplitN(str, []byte(":"), 2)

		if len(creds) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		username := string(creds[0])
		password := string(creds[1])

		// username is email
		requestID, _ := c.Get("Request-Id")
		user, err := m.userExtractor.ExtractUser(username, password, requestID.(string))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}

}

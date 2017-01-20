package util

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rachoac/service-skeleton-go/olio/common/models"
	jwt "gopkg.in/dgrijalva/jwt-go.v2"
)

func RequestToToken(c *gin.Context) *jwt.Token {
	token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
		signingKey := GetEnv("FAER_SIGNING_KEY", "secret")
		return []byte(signingKey), nil
	})
	if token == nil || !token.Valid || err != nil {
		return nil
	}

	return token
}

func TokenToEmail(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		signingKey := GetEnv("JWT_SIGNING_KEY", "secret")
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	subject := token.Claims["subject"].(string)
	return subject, nil
}

func TokenToClaims(token *jwt.Token) map[string]interface{} {
	var claims map[string]interface{} = make(map[string]interface{})
	for tokenKey := range token.Claims {
		value, ok := token.Claims[tokenKey]
		if ok {
			claims[tokenKey] = value
		}
	}

	return claims
}

func TokenToPermissions(tokenString string) []*models.Permission {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		signingKey := GetEnv("JWT_SIGNING_KEY", "secret")
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil
	}

	if token.Claims["permissions"] == nil {
		return nil
	}

	var permissionsToReturn []*models.Permission = []*models.Permission{}
	var permissionsStr string = token.Claims["permissions"].(string)
	permissions := strings.Split(permissionsStr, ",")
	for _, permission := range permissions {
		permissionParts := strings.Split(permission, "|")
		permissionType := permissionParts[0]
		permissionOperation := permissionParts[1]
		permissionObjectType := permissionParts[2]
		permissionObjectID := permissionParts[3]

		permissionsToReturn = append(permissionsToReturn, &models.Permission{
			Type:       permissionType,
			Operation:  permissionOperation,
			ObjectType: permissionObjectType,
			ObjectID:   permissionObjectID,
		})
	}

	return permissionsToReturn
}

func CreateToken(subject string, claims map[string]interface{}) string {
	signingKey := GetEnv("JWT_SIGNING_KEY", "secret")
	token := jwt.New(jwt.SigningMethodHS256)

	// Set a header and a claim
	token.Header["typ"] = "JWT"
	token.Claims["subject"] = subject

	// long expiration by default; can be overridden by claims map
	token.Claims["exp"] = time.Now().UTC().Add(time.Hour * 99999).Unix()

	if claims != nil {
		for claimKey := range claims {
			token.Claims[claimKey] = claims[claimKey]
		}
	}

	t, _ := token.SignedString([]byte(signingKey))

	return t
}

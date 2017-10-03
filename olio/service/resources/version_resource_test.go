package resources_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/olioapps/service-skeleton-go/olio/service/resources"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	router := gin.New()

	tt := []struct {
		name string
	}{
		{name: "Empty extractor"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			router.GET("/api/version", resources.NewVersionResource().GetVersion)
			req, _ := http.NewRequest("GET", "/api/version", nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, res.Body.String(), "{\n    \"serviceFrameworkVersion\": \"1.0.2\",\n    \"appVersion\": \"no version given\"\n}")
		})
	}
}

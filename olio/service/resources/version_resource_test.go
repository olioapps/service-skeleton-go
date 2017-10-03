package resources_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/olioapps/service-skeleton-go/olio/service/resources"
	"github.com/stretchr/testify/assert"
)

type VersionExtractor struct {
}

func NewVersionExtractor() *VersionExtractor {
	return &VersionExtractor{}
}

func (ve VersionExtractor) GetVersion() string {
	return "0.0.11"
}

func (ve VersionExtractor) GetAppName() string {
	return "cx-messaging"
}

func TestVersion(t *testing.T) {
	versionExtractor := NewVersionExtractor()

	tt := []struct {
		name string
	}{
		{name: "No extractor"},
		{name: "With extractor"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "No extractor" {
				router := gin.New()
				versionResource := resources.NewVersionResource()
				router.GET("/api/version", versionResource.GetVersion)
				req, _ := http.NewRequest("GET", "/api/version", nil)
				res := httptest.NewRecorder()
				router.ServeHTTP(res, req)
				assert.Equal(t, "{\n    \"serviceFrameworkVersion\": \""+resources.VERSION+"\",\n    \"appVersion\": \"no version given\"\n}", res.Body.String())
			}

			if tc.name == "With extractor" {
				router := gin.New()
				versionResource := resources.NewVersionResource()
				versionResource.AddVersionExtractor(versionExtractor)
				router.GET("/api/version", versionResource.GetVersion)
				req, _ := http.NewRequest("GET", "/api/version", nil)
				res := httptest.NewRecorder()
				router.ServeHTTP(res, req)
				assert.Equal(t, "{\n    \"serviceFrameworkVersion\": \""+resources.VERSION+"\",\n    \"appVersion\": \"cx-messaging-0.0.11\"\n}", res.Body.String())
			}
		})
	}
}

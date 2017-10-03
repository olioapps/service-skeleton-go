package resources

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Extractor struct {
}

func NewExtractor() Extractor {
	return Extractor{}
}

func (e Extractor) GetVersion() string {
	return "0.0.11"
}

func (e Extractor) GetAppName() string {
	return "cx-messaging"
}

func TestVersion(t *testing.T) {
	versionExtractor := NewExtractor()

	tt := []struct {
		name            string
		expectedVersion string
	}{
		{name: "No extractor", expectedVersion: "no version given"},
		{name: "With extractor", expectedVersion: "cx-messaging-0.0.11"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.New()
			versionResource := NewVersionResource()
			if tc.name == "With extractor" {
				versionResource.AddVersionExtractor(versionExtractor)
			}
			router.GET("/api/version", versionResource.getVersion)
			req, _ := http.NewRequest("GET", "/api/version", nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, "{\n    \"serviceFrameworkVersion\": \""+VERSION+"\",\n    \"appVersion\": \""+tc.expectedVersion+"\"\n}", res.Body.String())
		})
	}
}

package resources

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type Version struct {
	SkeletonVersion string `json:"skeletonVersion"`
	AppVersion      string `json:"appVersion"`
}

type VersionExtractor interface {
	GetVersion() string
}

type VersionResource struct {
	versionExtractor VersionExtractor
}

const VERSION = "0.0.1"

func NewVersionResource() *VersionResource {
	obj := VersionResource{}

	return &obj
}

func (resource *VersionResource) AddVersionExtractor(versionExtractor VersionExtractor) {
	resource.versionExtractor = versionExtractor
}

func (resource *VersionResource) Init(r *gin.Engine) {
	log.Debug("Setting up version resource.")

	r.GET("/api/version", resource.getVersion)
}

func (resource *VersionResource) getVersion(c *gin.Context) {
	skeletonVersion := VERSION
	var appVersion string
	if resource.versionExtractor != nil {
		appVersion = resource.versionExtractor.GetVersion()
	} else {
		appVersion = "no version given"
	}

	version := Version{
		AppVersion:      appVersion,
		SkeletonVersion: skeletonVersion,
	}

	w := c.Writer
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

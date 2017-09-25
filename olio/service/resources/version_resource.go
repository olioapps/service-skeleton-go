package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/thedataguild/faer/service/middleware"
)

type Version struct {
	skeletonVersion string `json:"skeletonVersion"`
	appVersion      string `json:"appVersion"`
}

type VersionExtractor interface {
	GetVersion()
}

type VersionResource struct {
	versionExtractor VersionExtractor
}

func NewVersionResource() VersionResource {
	obj := VersionResource{}

	return obj
}

func (resource VersionResource) AddVersionExtractor(versionExtractor VersionExtractor) {
	resource.versionExtractor = versionExtractor
}

func (resource VersionResource) init(r *gin.Engine) {
	lob.Debug("Setting up version resource.")

	r.GET("/api/version", resource.getVersion)
}

func (resource VersionResource) getVersion() {
	skeletonVersion := VERSION
	var appVersion string
	if resource.versionExtractor {
		appVersion = resource.versionExtractor.GetVersion()
	} else {
		appVersion = "no version given"
	}

	version := Version{
		appVersion: appVersion,
		skeletonVersion: skeletonVersion
	}

	resource.ReturnJSON(c, 200, version)
}

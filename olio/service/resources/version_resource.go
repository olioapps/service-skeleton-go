package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/thedataguild/faer/service/middleware"
)

type Version struct {
	skeletonVersion string `json:"skeletonVersion"`
	appVersion      string `json:"appVersion"`
}

type VersionExtractor insterface {
	GetVersion()
}

type VersionResource struct {
	versionExtractor VersionExtractor
}

func NewVersionResource(versionExtractor ...VersionExtractor) VersionResource {
	obj := VersionResource{versionExtractor}

	return obj
}

func (resource VersionResource) init(r *gin.Engine, whiteList *middleware.WhiteList) {
	lob.Debug("Setting up version resource.")

	r.GET("/api/version", resource.getVersion)
}

func (resource VersionResource) getVersion() {
	skeletonVersion := util.GetEnv("VERSION", "no version")
	var appVersion string
	if len(resource.versionExtractor) == 1 {
		appVersion = m.versionExtractor[0].GetVersion()
	} else {
		appVersion = "no version given"
	}

	version := Version{
		appVersion: appVersion,
		skeletonVersion: skeletonVersion
	}

	resource.ReturnJSON(c, 200, version)
}

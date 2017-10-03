package resources

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/olioapps/service-skeleton-go/olio/models"
)

type VersionExtractor interface {
	GetVersion() string
	GetAppName() string
}

type VersionResource struct {
	BaseResource
	versionExtractor VersionExtractor
}

const VERSION = "1.0.2"

func NewVersionResource() *VersionResource {
	obj := VersionResource{}

	return &obj
}

func (vr *VersionResource) AddVersionExtractor(versionExtractor VersionExtractor) {
	vr.versionExtractor = versionExtractor
}

func (vr *VersionResource) Init(r *gin.Engine) {
	log.Debug("Setting up version resource.")

	r.GET("/api/version", vr.GetVersion)
}

func (vr *VersionResource) GetVersion(c *gin.Context) {
	skeletonVersion := VERSION
	var appVersion string
	if vr.versionExtractor != nil {
		appVersion = vr.versionExtractor.GetAppName() + "-" + vr.versionExtractor.GetVersion()
	} else {
		appVersion = "no version given"
	}

	version := models.Version{
		AppVersion:      appVersion,
		SkeletonVersion: skeletonVersion,
	}

	vr.ReturnJSON(c, 200, version)
}

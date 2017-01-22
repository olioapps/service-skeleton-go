package resources

import (
	"bitbucket.com/olioapps/service-skeleton-go/olio/service/middleware"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type PingResource struct {
}

func NewPingResource() PingResource {
	obj := PingResource{}

	return obj
}

func (resource PingResource) Init(r *gin.Engine, whiteList *middleware.WhiteList) {
	log.Debug("Setting up ping resource.")

	r.GET("/ping", resource.ping)
}

func (resource PingResource) ping(c *gin.Context) {
	c.Writer.WriteString("pong")
}

package resources

import (
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/olioapps/service-skeleton-go/olio/service/middleware"
)

type PingResource struct {
}

func NewPingResource() PingResource {
	obj := PingResource{}

	return obj
}

func (resource PingResource) Init(r *gin.Engine, whiteList *middleware.WhiteList) {
	log.Debug("Setting up ping resource.")

	r.GET("/api/ping", resource.ping)
}

func (resource PingResource) ping(c *gin.Context) {
	c.Writer.WriteString("pong")
}

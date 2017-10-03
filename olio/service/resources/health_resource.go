package resources

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olioapps/service-skeleton-go/olio/models"
	"github.com/siddontang/go/log"
)

type HealthResource struct {
	BaseResource
	uptimeExtractor UptimeExtractor
}

type UptimeExtractor interface {
	GetUptime() time.Duration
}

func NewHealthResource() *HealthResource {
	obj := HealthResource{}
	return &obj
}

func (hr *HealthResource) Init(e *gin.Engine) {
	log.Debug("setting up health resource")

	e.GET("/api/health", hr.getHealth)
}

func (hr *HealthResource) AddUptimeExtractor(uptimeExtractor UptimeExtractor) {
	hr.uptimeExtractor = uptimeExtractor
}

func (hr *HealthResource) getHealth(c *gin.Context) {
	var uptime string
	if hr.uptimeExtractor != nil {
		tmp := int(hr.uptimeExtractor.GetUptime())
		uptime = fmt.Sprintf("%.3f", float64(tmp)/(1000*60*60)) + " hours"
	}

	health := models.Health{
		Uptime: uptime,
	}

	w := c.Writer
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(health)
	if err != nil {
		hr.ReturnError(c, 500, err.Error())
		return
	}

	w.Write(js)
}

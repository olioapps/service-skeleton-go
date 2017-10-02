package resources

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	olioMiddleware "github.com/olioapps/service-skeleton-go/olio/service/middleware"
	"github.com/siddontang/go/log"
)

type HealthResource struct {
	uptimeExtractor UptimeExtractor
}

type UptimeExtractor interface {
	GetUptime() time.Duration
}

// should go in models
type Health struct {
	Uptime time.Duration `json:"uptime"`
	// DataStorePingSuccess bool          `json:"dataStorePingSuccess"`
}

func NewHealthResource() *HealthResource {
	obj := HealthResource{}
	return &obj
}

func (hr *HealthResource) Init(e *gin.Engine, whiteList *olioMiddleware.WhiteList) {
	log.Debug("setting up health resource")

	e.GET("/api/health", hr.getHealth)
}

func (hr *HealthResource) AddUptimeExtractor(uptimeExtractor UptimeExtractor) {
	hr.uptimeExtractor = uptimeExtractor
}

func (hr *HealthResource) getHealth(c *gin.Context) {
	var uptime time.Duration
	if hr.uptimeExtractor != nil {
		uptime = hr.uptimeExtractor.GetUptime()
	}

	health := Health{
		Uptime: uptime,
	}

	w := c.Writer
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(health)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)

}

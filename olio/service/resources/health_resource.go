package resources

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/olioapps/service-skeleton-go/olio/dao"
	"github.com/olioapps/service-skeleton-go/olio/extractors"
	"github.com/olioapps/service-skeleton-go/olio/models"
	log "github.com/sirupsen/logrus"
)

type HealthResource struct {
	BaseResource
	uptimeExtractor   extractors.UptimeExtractor
	dbHealthExtractor extractors.DbHealthExtractor
}

func NewHealthResource() *HealthResource {
	obj := HealthResource{}
	return &obj
}

func (hr *HealthResource) Init(e *gin.Engine) {
	log.Debug("setting up health resource")

	e.GET("/api/health", hr.getHealth)
}

func (hr *HealthResource) AddUptimeExtractor(uptimeExtractor extractors.UptimeExtractor) {
	hr.uptimeExtractor = uptimeExtractor
}

func (hr *HealthResource) AddDbHealthExtractor(dbHealthExtractor extractors.DbHealthExtractor) {
	hr.dbHealthExtractor = dbHealthExtractor
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

	if hr.dbHealthExtractor != nil {
		connectionManager := dao.NewConnectionManager(hr.dbHealthExtractor.GetDbExtractor())
		if err := connectionManager.Ping(); err != nil {
			health.DbOk = false
			log.Error("Database not ok: ", err)
		} else {
			health.DbOk = true
		}

		err := connectionManager.Close()
		if err != nil {
			log.Error("Error closing db: ", err)
		}
	}

	hr.ReturnJSON(c, 200, health)
}

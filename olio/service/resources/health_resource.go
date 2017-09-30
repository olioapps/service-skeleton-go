package resources

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	olioMiddleware "github.com/olioapps/service-skeleton-go/olio/service/middleware"
	"github.com/siddontang/go/log"
)

type HealthResource struct {
}

type Health struct {
	pingSuccess          bool
	uptime               time.Time
	dataStorePingSuccess bool
}

func NewHealthResource() *HealthResource {
	obj := HealthResource{}
	return &obj
}

func (hr HealthResource) Init(e *gin.Engine, whiteList *olioMiddleware.WhiteList) {
	log.Debug("setting up health resource")

	e.GET("/api/health", hr.getHealth)
}

func (hr HealthResource) getHealth(c *gin.Context) {
	// var pingSuccess bool

	w := c.Writer
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://cx-messaging.herokuapp.com/api/ping")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBodyText, _ := ioutil.ReadAll(resp.Body)
	if string(respBodyText) == "pong" {
		c.Writer.WriteString("success")
	}

	c.Writer.WriteString("fail")
}

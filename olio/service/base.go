package service

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	network "github.com/glibs/gin-webserver"
	olioMiddleware "github.com/olioapps/service-skeleton-go/olio/service/middleware"
	olioResources "github.com/olioapps/service-skeleton-go/olio/service/resources"
	"github.com/olioapps/service-skeleton-go/olio/util"
)

type OlioDaemon interface {
	Start()
	Stop()
}

type OlioResourceHandler interface {
	Init(*gin.Engine, *olioMiddleware.WhiteList)
}

type OlioBaseService struct {
	GinEngine       *gin.Engine
	server          *network.WebServer
	daemons         []OlioDaemon
	versionResource map[string]*olioResources.VersionResource
	healthResource  map[string]*olioResources.HealthResource
}

func New() *OlioBaseService {
	service := OlioBaseService{}
	service.GinEngine = gin.Default()
	service.versionResource = make(map[string]*olioResources.VersionResource)
	service.healthResource = make(map[string]*olioResources.HealthResource)

	return &service
}

func (obs *OlioBaseService) Init(whitelist *olioMiddleware.WhiteList, middlewares []gin.HandlerFunc, resources []OlioResourceHandler) {
	log.Info("Initializing RESTful service.")

	log.Debug("Setting up middleware.")

	obs.GinEngine.Use(whitelist.Handler)
	obs.GinEngine.Use(olioMiddleware.RequestId)

	for _, middleware := range middlewares {
		obs.GinEngine.Use(middleware)
	}

	healthResource := olioResources.NewHealthResource()
	obs.healthResource["health"] = healthResource
	healthResource.Init(obs.GinEngine, whitelist)

	versionResource := olioResources.NewVersionResource()
	obs.versionResource["version"] = versionResource
	versionResource.Init(obs.GinEngine)

	pingResource := olioResources.NewPingResource()
	pingResource.Init(obs.GinEngine, whitelist)
	for _, resource := range resources {
		resource.Init(obs.GinEngine, whitelist)
	}

	log.Debug("Setting up routes.")
}

func (obs *OlioBaseService) AddDaemon(daemon OlioDaemon) {
	obs.daemons = append(obs.daemons, daemon)
}

func (obs *OlioBaseService) AddVersionProvider(versionExtractor olioResources.VersionExtractor) {
	versionResource := obs.versionResource["version"]
	if versionResource != nil {
		versionResource.AddVersionExtractor(versionExtractor)
	}
}

func (obs *OlioBaseService) AddUptimeProvider(uptimeExtractor olioResources.UptimeExtractor) {
	healthResource := obs.healthResource["health"]
	if healthResource != nil {
		healthResource.AddUptimeExtractor(uptimeExtractor)
	}
}

func (obs *OlioBaseService) Start() {
	for _, daemon := range obs.daemons {
		daemon.Start()
	}

	servicePort := util.GetEnv("PORT", "9090")
	host := ":" + servicePort
	obs.server = network.InitializeWebServer(obs.GinEngine, host)
	obs.server.Start()
}

func (obs *OlioBaseService) Stop() {
	for _, daemon := range obs.daemons {
		daemon.Stop()
	}
	obs.server.Stop()
}

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
	GinEngine *gin.Engine
	server    *network.WebServer
	daemons   []OlioDaemon
}

func New() *OlioBaseService {
	service := OlioBaseService{}
	service.GinEngine = gin.Default()

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

	pingResource := olioResources.NewPingResource()
	pingResource.Init(obs.GinEngine, whitelist)
	for _, resource := range resources {
		resource.Init(obs.GinEngine, whitelist)
	}

	log.Debug("Setting up routes.")
}

func (service *OlioBaseService) AddDaemon(daemon OlioDaemon) {
	service.daemons = append(service.daemons, daemon)
}

func (service *OlioBaseService) Start() {
	for _, daemon := range service.daemons {
		daemon.Start()
	}

	servicePort := util.GetEnv("SERVICE_PORT", "9090")
	host := ":" + servicePort
	service.server = network.InitializeWebServer(service.GinEngine, host)
	service.server.Start()
}

func (service *OlioBaseService) Stop() {
	for _, daemon := range service.daemons {
		daemon.Stop()
	}
	service.server.Stop()
}

func (service *OlioBaseService) AddVersionProvider(versionExtractor ...VersionExtractor) {
	versionResource := olioResources.NewVersionResource(versionExtractor)
	versionResource.Init(obs.GinEngine)
}

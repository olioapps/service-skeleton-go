package service

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	network "github.com/glibs/gin-webserver"
	olioMiddleware "github.com/olioapps/service-skeleton-go/olio/service/middleware"
	olioResources "github.com/olioapps/service-skeleton-go/olio/service/resources"
	"github.com/olioapps/service-skeleton-go/olio/util"
	"github.com/rachoac/service-skeleton-go/olio/service/resources"
)

type OlioDaemon interface {
	Start()
	Stop()
}

type OlioResourceHandler interface {
	Init(*gin.Engine, *olioMiddleware.WhiteList)
}

type OlioBaseService struct {
	GinEngine     *gin.Engine
	server        *network.WebServer
	daemons       []OlioDaemon
	coreResources map[string]*resources.VersionResource
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

	versionResource := olioResources.NewVersionResource()
	obs.coreResources["version"] = &versionResource
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

func (obs *OlioBaseService) AddVersionProvider(versionExtractor resources.VersionExtractor) {
	versionResource := obs.coreResources["version"]
	if versionResource != nil {
		versionResource.AddVersionExtractor(versionExtractor)
	}
}

func (obs *OlioBaseService) Start() {
	for _, daemon := range obs.daemons {
		daemon.Start()
	}

	servicePort := util.GetEnv("SERVICE_PORT", "9090")
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

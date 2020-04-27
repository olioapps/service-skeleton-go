package api

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/olioapps/service-skeleton-go/olio/dao"
	"github.com/olioapps/service-skeleton-go/olio/db"
)

type OlioBaseCoreAPI struct {
	ConnectionManager *dao.ConnectionManager
}

func (api *OlioBaseCoreAPI) Init() {
	log.Info("Initializing core api.")
}

func (api *OlioBaseCoreAPI) RunMigrations(migrations []db.Migration) {
	migrationManager := db.NewMigrationManager(api.ConnectionManager, migrations)
	if err := migrationManager.Migrate(); err != nil {
		log.Fatal("Failed to run migrations: ", err)
		os.Exit(1)
	}
}
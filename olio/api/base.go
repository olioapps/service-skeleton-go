package api

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/rachoac/service-skeleton-go/olio/common/filters"
	"github.com/rachoac/service-skeleton-go/olio/common/models"
	"github.com/rachoac/service-skeleton-go/olio/dao"
	"github.com/rachoac/service-skeleton-go/olio/db"
)

type OlioAPI interface {
	GetAPIName() string
	Create(models.AccessContext, models.OlioModel) error
	Update(models.AccessContext, models.OlioModel) error
	Delete(models.AccessContext, string) error
	Get(models.AccessContext, string) (models.OlioModel, error)
	Find(models.AccessContext, filters.OlioFilter) ([]models.OlioModel, error)
}

type CoreAPI struct {
	ConnectionManager *dao.ConnectionManager
}

func NewCoreAPI(connectionManager *dao.ConnectionManager) *CoreAPI {
	api := CoreAPI{}

	api.ConnectionManager = dao.NewConnectionManager()

	return &api
}

func (api *CoreAPI) RunMigrations(migrations []db.Migration) {
	migrationManager := db.NewMigrationManager(api.ConnectionManager, migrations)
	if err := migrationManager.Migrate(); err != nil {
		log.Fatal("Failed to run migrations: ", err)
		os.Exit(1)
	}

}

func (api *CoreAPI) RegisterOlioAPI(olioAPI OlioAPI) error {
	log.Info("Registering API " + olioAPI.GetAPIName())
	return nil
}

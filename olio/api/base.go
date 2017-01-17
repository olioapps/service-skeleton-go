package api

import (
	"log"
	"os"

	"github.com/rachoac/service-skeleton-go/olio/dao"
	"github.com/rachoac/service-skeleton-go/olio/db"
)

//type IDAware interface {
//	GetID() string
//}
//
//type API interface {
//	GetByID()
//}
//
//type AccessContext interface {
//}
//
//type Model interface {
//	GetID() string
//	GetType() string
//	GetOwnerID() string
//	GetCreatedAt() time.Time
//	GetUpdatedAt() time.Time
//}
//
//type User interface {
//}

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

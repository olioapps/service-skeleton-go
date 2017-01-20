package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/rachoac/service-skeleton-go/olio/common/filters"
	"github.com/rachoac/service-skeleton-go/olio/common/models"
)

type OlioDAO interface {
	Insert(models.OlioModel) error
	Update(models.OlioModel) error
	Find(filters.OlioFilter) ([]models.OlioModel, error)
	Delete(id string) error
}

type BaseDAO struct {
	ConnectionProvider ConnectionProvider
}

func (d *BaseDAO) Db() *gorm.DB {
	return d.ConnectionProvider.GetDb()
}

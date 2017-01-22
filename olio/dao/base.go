package dao

import (
	"bitbucket.com/olioapps/service-skeleton-go/olio/common/filters"
	"bitbucket.com/olioapps/service-skeleton-go/olio/common/models"
	"github.com/jinzhu/gorm"
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

package dao

import (
	"github.com/rachoac/service-skeleton-go/olio/common/models"
	"github.com/rachoac/service-skeleton-go/olio/common/filter"
)

type PermissionsDAO struct {
	BaseDAO
}

func NewPermissionsDAO(connectionManager ConnectionProvider) *PermissionsDAO {
	dao := PermissionsDAO{
		BaseDAO{connectionManager, models.AccessToken{}},
	}

	return &dao
}

func (self *PermissionsDAO) Find(filter *filters.AccessTokenFilter) ([]models.AccessToken, error) {
	db := self.connectionManager.GetDb()

	if filter.Token != "" {
		db = db.Where("token = ?", filter.Token)
	}

	results := []models.AccessToken{}
	db = db.Find(&results)

	return results, db.Error
}

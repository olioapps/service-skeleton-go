package dao

import (
	"github.com/olioapps/service-skeleton-go/olio/common/filter"
	"github.com/olioapps/service-skeleton-go/olio/common/models"
)

type PermissionsDAO struct {
	StringBaseDAO
}

func NewPermissionsDAO(connectionManager ConnectionProvider) *PermissionsDAO {
	dao := PermissionsDAO{
		StringBaseDAO{
			BaseDAO{connectionManager, models.AccessToken{}},
		},
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

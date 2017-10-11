package dao

import (
	"github.com/jinzhu/gorm"
)

type IDAware interface {
	GetID() interface{}
	SetID(id interface{})
}

type BaseDAO struct {
	ConnectionManager ConnectionProvider
	Model             interface{}
}

type DAO interface {
	Delete(object IDAware, tx ...*gorm.DB) error
	DeleteByID(id interface{})
	Insert(object IDAware) error
	Update(object IDAware) error
	GetConnectionManager() ConnectionProvider
}

func (d *BaseDAO) Db() *gorm.DB {
	return d.ConnectionManager.GetDb()
}

func (d *BaseDAO) GetConnectionManager() ConnectionProvider {
	return d.ConnectionManager
}

func (d *BaseDAO) Delete(object IDAware, tx ...*gorm.DB) error {
	var db *gorm.DB = nil

	hasTransaction := len(tx) > 0
	if hasTransaction {
		db = tx[0]
	} else {
		db = d.ConnectionManager.GetDb()
	}

	if err := db.Delete(object).Error; err != nil {
		return err
	}
	return db.Error
}

func (d *BaseDAO) DeleteByID(id interface{}) error {
	db := d.ConnectionManager.GetDb()
	db = db.Where("id = ?", id).Delete(d.Model)
	return db.Error
}

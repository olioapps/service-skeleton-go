package dao

import (
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/olioapps/service-skeleton-go/olio/util"
)

type IDAware interface {
	GetID() string
	SetID(id string)
}

type BaseDAO struct {
	connectionManager ConnectionProvider
	model interface{}
}

func (d *BaseDAO) Db() *gorm.DB {
	return d.connectionManager.GetDb()
}

func (d *BaseDAO) Delete(object IDAware, tx ...*gorm.DB) error {
	var db *gorm.DB = nil

	hasTransaction := len(tx) > 0
	if hasTransaction {
		db = tx[0]
	} else {
		db = d.connectionManager.GetDb()
	}

	if err := db.Delete(object).Error; err != nil {
		return err
	}
	return db.Error
}

func (d *BaseDAO) DeleteByID(id string) error {
	db := d.connectionManager.GetDb()
	db = db.Where("id = ?", id).Delete(d.model)
	return db.Error
}

func (d *BaseDAO) Insert(object IDAware) error {
	if object.GetID() != "" {
		return errors.New("Cannot insert an object that already has an ID")
	}
	object.SetID(util.RandomString())
	db := d.connectionManager.GetDb()
	return db.Create(object).Error
}

func (d *BaseDAO) Update(object IDAware) error {
	if object.GetID() == "" {
		return errors.New("Cannot update object without an ID")
	}
	db := d.connectionManager.GetDb()
	return db.Save(object).Error
}

package dao

import (
	"errors"

	"github.com/olioapps/service-skeleton-go/olio/util"
)

type StringBaseDAO struct {
	BaseDAO
}

func (d *StringBaseDAO) Insert(object IDAware) error {
	id := object.GetID().(string)
	if id != "" {
		return errors.New("Cannot insert an object that already has an ID")
	}
	object.SetID(util.RandomString())
	db := d.connectionManager.GetDb()
	return db.Create(object).Error
}

func (d *StringBaseDAO) Update(object IDAware) error {
	id := object.GetID().(string)
	if id == "" {
		return errors.New("Cannot update object without an ID")
	}
	db := d.connectionManager.GetDb()
	return db.Save(object).Error
}

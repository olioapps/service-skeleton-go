package dao

import (
	"errors"
)

type NumericBaseDAO struct {
	BaseDAO
}

func (d *NumericBaseDAO) Insert(object IDAware) error {
	id := object.GetID().(int64)

	if id > 0 {
		return errors.New("Cannot insert an object that already has an ID")
	}
	db := d.connectionManager.GetDb()
	return db.Create(object).Error
}

func (d *NumericBaseDAO) Update(object IDAware) error {
	id := object.GetID().(int64)
	if id < 1 {
		return errors.New("Cannot update object without an ID")
	}
	db := d.connectionManager.GetDb()
	return db.Save(object).Error
}

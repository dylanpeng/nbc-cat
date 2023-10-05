package model

import (
	"cat/common"
	"cat/common/entity"
	"errors"
	"gorm.io/gorm"
)

type baseDBModel struct {
	readInstance  string
	writeInstance string
	rwSplite      bool
}

var ErrPrimaryAttrEmpty = errors.New("primary attribute is empty")

func createDBModel(readInstance, writeInstance string, rwSplite bool) *baseDBModel {
	return &baseDBModel{
		readInstance:  readInstance,
		writeInstance: writeInstance,
		rwSplite:      rwSplite,
	}
}

func (m *baseDBModel) getDB(write bool) (*gorm.DB, error) {
	if m.rwSplite && !write {
		return common.GetDB(m.readInstance)
	}

	return common.GetDB(m.writeInstance)
}

func (m *baseDBModel) Add(e entity.IEntity) error {
	db, err := m.getDB(true)

	if err != nil {
		return err
	}

	return db.Create(e).Error
}

func (m *baseDBModel) Get(e entity.IEntity) (exist bool, err error) {
	if !e.PrimarySeted() {
		err = ErrPrimaryAttrEmpty
		return
	}

	db, err := m.getDB(false)

	if err != nil {
		return false, err
	}

	err = db.First(e).Error

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *baseDBModel) Update(e entity.IEntity, props map[string]interface{}) error {
	if !e.PrimarySeted() {
		return ErrPrimaryAttrEmpty
	}

	db, err := m.getDB(true)

	if err != nil {
		return err
	}

	if props == nil {
		return db.Save(e).Error
	} else {
		return db.Model(e).Updates(props).Error
	}
}

func (m *baseDBModel) Remove(e entity.IEntity) error {
	if !e.PrimarySeted() {
		return ErrPrimaryAttrEmpty
	}

	db, err := m.getDB(true)

	if err != nil {
		return err
	}

	return db.Delete(e).Error
}

func (m *baseDBModel) RemoveByIds(e entity.IEntity, ids []int64) (err error) {
	db, err := m.getDB(true)

	if err != nil {
		return
	}

	return db.Delete(e, "id IN (?)", ids).Error
}

func (m *baseDBModel) GetIdsByParams(e entity.IEntity, where string, params ...interface{}) (ids []int64, err error) {
	db, err := m.getDB(false)

	if err != nil {
		return
	}

	rows, err := db.Table(e.TableName()).Select("id").Where(where, params...).Rows()

	if err != nil {
		return
	}

	defer func() { _ = rows.Close() }()

	ids = make([]int64, 0, 8)

	for rows.Next() {
		var id int64

		if e := rows.Scan(&id); e != nil {
			continue
		}

		ids = append(ids, id)
	}

	return
}

func (m *baseDBModel) IncrNum(e entity.IEntity, field string, value int64) (err error) {
	if !e.PrimarySeted() {
		return ErrPrimaryAttrEmpty
	}

	db, err := m.getDB(true)

	if err != nil {
		return
	}

	return db.Model(e).Update(field, gorm.Expr(field+"+?", value)).Error
}

func (m *baseDBModel) Raw(write bool, sql string, values ...interface{}) (db *gorm.DB, err error) {
	db, err = m.getDB(write)

	if err != nil {
		return
	}

	db = db.Raw(sql, values...)
	err = db.Error
	return
}

func (m *baseDBModel) Exec(write bool, sql string, values ...interface{}) (db *gorm.DB, err error) {
	db, err = m.getDB(write)

	if err != nil {
		return
	}

	db = db.Exec(sql, values...)
	err = db.Error
	return
}

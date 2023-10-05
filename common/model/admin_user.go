package model

import (
	"cat/common/entity"
	"gorm.io/gorm"
	"time"
)

var AdminUser = &adminUserModel{
	baseModel: createModel(
		"main-slave",
		"main-master",
		"main-slave",
		"main-master",
		false,
		"cat:admin:user",
		time.Minute*10,
	),
}

type adminUserModel struct {
	*baseModel
}

func (m *adminUserModel) GetByName(name string) (item *entity.DataAdminUser, err error) {
	db, err := m.DB.getDB(false)
	if err != nil {
		return
	}

	oItem := &entity.DataAdminUser{}
	err = db.Where("`name`=?", name).First(oItem).Error

	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	}

	if err != nil {
		return
	}

	item = oItem
	return
}

func (m *adminUserModel) List(userId, roleId int64, userName, email string, status, page, pageSize int) (items []*entity.DataAdminUser, total int64, err error) {
	db, err := m.DB.getDB(false)

	if err != nil {
		return
	}

	if userId > 0 {
		db = db.Where("`id`=?", userId)
	}

	if roleId > 0 {
		db = db.Where("JSON_CONTAINS(`roles`, ?, '$')", roleId)
	}

	if userName != "" {
		db = db.Where("`name`=?", userName)
	}

	if email != "" {
		db = db.Where("`email`=?", email)
	}

	if status > 0 {
		db = db.Where("`status`=?", status)
	}

	err = db.Model(&entity.DataAdminUser{}).Count(&total).Error

	if err != nil {
		return
	}

	db = db.Order("`id` DESC")

	if pageSize > 0 {
		offset := (page - 1) * pageSize
		db = db.Limit(pageSize).Offset(offset)
	}

	oItems := make([]*entity.DataAdminUser, 0, pageSize)
	err = db.Find(&oItems).Error

	if err != nil {
		return
	}

	items = oItems
	return
}

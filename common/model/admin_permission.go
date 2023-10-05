package model

import (
	"cat/common/entity"
	"time"
)

var AdminPermission = &adminPermissionModel{
	baseModel: createModel(
		"admin-slave",
		"admin-master",
		"main-slave",
		"main-master",
		false,
		"cat:admin:permission",
		time.Minute*10,
	),
}

type adminPermissionModel struct {
	*baseModel
}

func (m *adminRoleModel) GetAllKey() string {
	return m.Cache.GetKey("all")
}

func (m *adminRoleModel) GetAll() (items []*entity.DataAdminPermission, err error) {
	items = make([]*entity.DataAdminPermission, 0)
	key := m.GetAllKey()

	exist, err := m.Cache.Get(key, &items)
	if exist {
		return
	}

	db, err := m.DB.getDB(false)
	if err != nil {
		return
	}

	items = make([]*entity.DataAdminPermission, 0)
	err = db.Find(&items).Error
	if err != nil {
		return
	}

	if len(items) > 0 {
		m.Cache.Set(key, items)
	}

	return
}

package model

import (
	"cat/common/entity"
	"time"
)

var AdminPermissionPath = &adminPermissionPathModel{
	baseModel: createModel(
		"admin-slave",
		"admin-master",
		"main-slave",
		"main-master",
		false,
		"cat:admin:permission:path",
		time.Minute*10,
	),
}

type adminPermissionPathModel struct {
	*baseModel
}

func (m *adminPermissionPathModel) GetPathBySymbol(symbols []string) (items []*entity.DataAdminPermissionPath, err error) {
	items = make([]*entity.DataAdminPermissionPath, 0)

	db, err := m.DB.getDB(false)
	if err != nil {
		return
	}

	err = db.Model(&entity.DataAdminPermissionPath{}).Where("perm_symbol in ?", symbols).Find(&items).Error
	if err != nil {
		return
	}

	return
}

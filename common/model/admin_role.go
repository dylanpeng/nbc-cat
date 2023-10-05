package model

import (
	"cat/common/entity"
	"time"
)

var AdminRole = &adminRoleModel{
	baseModel: createModel(
		"main-slave",
		"main-master",
		"main-slave",
		"main-master",
		false,
		"cat:admin:role",
		time.Minute*10,
	),
}

type adminRoleModel struct {
	*baseModel
}

func (m *adminRoleModel) QueryByIds(ids []int64) (items []*entity.DataAdminRole, err error) {
	db, err := m.DB.getDB(false)
	if err != nil {
		return
	}

	items = make([]*entity.DataAdminRole, 0)
	err = db.Where(ids).Find(&items).Error
	if err != nil {
		return
	}

	return
}

func (m *adminRoleModel) GetByIds(ids []int64) (items []*entity.DataAdminRole, err error) {
	oItems := make([]entity.IEntity, 0, len(ids))

	for _, id := range ids {
		oItems = append(oItems, &entity.DataAdminRole{ID: id})
	}

	items = make([]*entity.DataAdminRole, 0, len(ids))
	cachedItems, missItems := m.Cache.GetEntitys(oItems)

	for _, e := range cachedItems {
		items = append(items, e.(*entity.DataAdminRole))
	}

	missIds := make([]int64, 0, len(ids))

	for _, e := range missItems {
		missIds = append(missIds, e.(*entity.DataAdminRole).ID)
	}

	if len(missIds) == 0 {
		return
	}

	queryItems, err := m.QueryByIds(missIds)

	if err != nil {
		return nil, err
	}

	if len(queryItems) == 0 {
		return
	}

	items = append(items, queryItems...)
	cacheItems := make([]entity.IEntity, 0, len(queryItems))

	for _, v := range queryItems {
		cacheItems = append(cacheItems, v)
	}

	_ = m.Cache.SetEntitys(cacheItems)
	return
}

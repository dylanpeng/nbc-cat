package model

import (
	"cat/common/entity"
	"time"
)

var AdminPermission = &adminPermissionModel{
	baseModel: createModel(
		"main-slave",
		"main-master",
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

func (m *adminPermissionModel) GetAllKey() string {
	return m.Cache.GetKey("all")
}

func (m *adminPermissionModel) GetAll() (items []*entity.DataAdminPermission, err error) {
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

func (m *adminPermissionModel) QueryByIds(ids []int64) (items []*entity.DataAdminPermission, err error) {
	db, err := m.DB.getDB(false)
	if err != nil {
		return
	}

	items = make([]*entity.DataAdminPermission, 0)
	err = db.Where(ids).Find(&items).Error
	if err != nil {
		return
	}

	return
}

func (m *adminPermissionModel) GetByIds(ids []int64) (items []*entity.DataAdminPermission, err error) {
	oItems := make([]entity.IEntity, 0, len(ids))

	for _, id := range ids {
		oItems = append(oItems, &entity.DataAdminPermission{ID: id})
	}

	items = make([]*entity.DataAdminPermission, 0, len(ids))
	cachedItems, missItems := m.Cache.GetEntitys(oItems)

	for _, e := range cachedItems {
		items = append(items, e.(*entity.DataAdminPermission))
	}

	missIds := make([]int64, 0, len(ids))

	for _, e := range missItems {
		missIds = append(missIds, e.(*entity.DataAdminPermission).ID)
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

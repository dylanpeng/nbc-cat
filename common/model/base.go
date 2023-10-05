package model

import (
	"cat/common/entity"
	"time"
)

type baseModel struct {
	DB    *baseDBModel
	Cache *baseCacheModel
}

func createModel(dbReadInstance, dbWriteInstance, cacheReadInstance, cacheWriteInstance string, rwSplite bool, cachePrefix string, cacheExpire time.Duration) *baseModel {
	return &baseModel{
		DB:    createDBModel(dbReadInstance, dbWriteInstance, rwSplite),
		Cache: createCacheModel(cacheReadInstance, cacheWriteInstance, rwSplite, cachePrefix, cacheExpire),
	}
}

func (m *baseModel) Add(e entity.IEntity) error {
	err := m.DB.Add(e)

	if err == nil {
		_ = m.Cache.Set(m.Cache.PrimaryKey(e), e)
	}

	return err
}

func (m *baseModel) Get(e entity.IEntity) (exist bool, err error) {
	if !e.PrimarySeted() {
		err = ErrPrimaryAttrEmpty
		return
	}

	exist, err = m.Cache.Get(m.Cache.PrimaryKey(e), e)

	if exist && err == nil {
		return
	}

	exist, err = m.DB.Get(e)

	if exist {
		_ = m.Cache.Set(m.Cache.PrimaryKey(e), e)
	}

	return
}

func (m *baseModel) Update(e entity.IEntity, props map[string]interface{}) error {
	if !e.PrimarySeted() {
		return ErrPrimaryAttrEmpty
	}

	err := m.DB.Update(e, props)

	if err != nil {
		return err
	}

	return m.Cache.Del(m.Cache.PrimaryKey(e))
}

func (m *baseModel) Remove(e entity.IEntity) error {
	if !e.PrimarySeted() {
		return ErrPrimaryAttrEmpty
	}

	err := m.DB.Remove(e)

	if err != nil {
		return err
	}

	return m.Cache.Del(m.Cache.PrimaryKey(e))
}

func (m *baseModel) IncrNum(e entity.IEntity, field string, value int64) (err error) {
	if !e.PrimarySeted() {
		return ErrPrimaryAttrEmpty
	}

	err = m.DB.IncrNum(e, field, value)

	if err != nil {
		return
	}

	return m.Cache.Del(m.Cache.PrimaryKey(e))
}

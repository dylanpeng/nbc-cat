package model

import (
	"cat/common"
	"cat/common/entity"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type baseCacheModel struct {
	readInstance  string
	writeInstance string
	rwSplite      bool

	prefix string
	expire time.Duration
}

func createCacheModel(readInstance, writeInstance string, rwSplite bool, prefix string, expire time.Duration) *baseCacheModel {
	return &baseCacheModel{
		readInstance:  readInstance,
		writeInstance: writeInstance,
		rwSplite:      rwSplite,
		prefix:        prefix,
		expire:        expire,
	}
}

func (m *baseCacheModel) GetKey(items ...interface{}) string {
	return common.GetKey(m.prefix, items...)
}

func (m *baseCacheModel) PrimaryKey(e entity.IEntity) string {
	return m.GetKey(e.PrimaryPairs()...)
}

func (m *baseCacheModel) getCache(write bool) (*redis.Client, error) {
	if m.rwSplite && !write {
		return common.GetCache(m.readInstance)
	}

	return common.GetCache(m.writeInstance)
}

func (m *baseCacheModel) Get(key string, data interface{}) (bool, error) {
	cache, err := m.getCache(false)

	if err != nil {
		return false, err
	}

	result, err := cache.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(result), data)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *baseCacheModel) Set(key string, data interface{}, exp ...time.Duration) error {
	cache, err := m.getCache(true)

	if err != nil {
		return err
	}

	result, err := json.Marshal(data)

	if err != nil {
		return err
	}

	expire := m.expire

	if len(exp) > 0 {
		expire = exp[0]
	}

	return cache.Set(context.Background(), key, result, expire).Err()
}

func (m *baseCacheModel) Del(key string) error {
	cache, err := m.getCache(true)

	if err != nil {
		return err
	}

	return cache.Del(context.Background(), key).Err()
}

func (m *baseCacheModel) GetEntitys(entitys []entity.IEntity) (getEntitys []entity.IEntity, missEntitys []entity.IEntity) {
	cache, err := m.getCache(false)

	if err != nil {
		return nil, entitys
	}

	keys := make([]string, 0, len(entitys))

	for _, e := range entitys {
		keys = append(keys, m.PrimaryKey(e))
	}

	items, err := cache.MGet(context.Background(), keys...).Result()

	if err != nil {
		return nil, entitys
	}

	getEntitys, missEntitys = make([]entity.IEntity, 0, len(entitys)), make([]entity.IEntity, 0, len(entitys))

	if len(items) != len(entitys) {
		return nil, entitys
	}

	for index, e := range entitys {
		if item := items[index]; item != nil {
			if err := json.Unmarshal([]byte(item.(string)), e); err == nil {
				getEntitys = append(getEntitys, e)
				continue
			}
		}

		missEntitys = append(missEntitys, e)
	}

	return
}

func (m *baseCacheModel) DelEntitys(entitys []entity.IEntity) error {
	cache, err := m.getCache(true)

	if err != nil {
		return err
	}

	keys := make([]string, 0, len(entitys))

	for _, e := range entitys {
		keys = append(keys, m.PrimaryKey(e))
	}

	return cache.Del(context.Background(), keys...).Err()
}

func (m *baseCacheModel) SetEntitys(entitys []entity.IEntity) error {
	cache, err := m.getCache(true)

	if err != nil {
		return err
	}

	pipe := cache.Pipeline()

	for _, e := range entitys {
		if data, err := json.Marshal(e); err == nil {
			pipe.Set(context.Background(), m.PrimaryKey(e), data, m.expire)
		}
	}

	_, err = pipe.Exec(context.Background())
	return err
}

func (m *baseCacheModel) IncrBy(key string, value int64, expiration time.Duration) (res int64, err error) {
	cache, err := m.getCache(true)

	if err != nil {
		return
	}

	pipe := cache.Pipeline()
	pipe.IncrBy(context.Background(), key, value)
	pipe.Expire(context.Background(), key, expiration)
	cmds, err := pipe.Exec(context.Background())

	if err != nil {
		return
	}

	if len(cmds) == 2 {
		if cmd, ok := cmds[0].(*redis.IntCmd); ok {
			res, err = cmd.Result()

			if err == redis.Nil {
				err = nil
			}
		}
	}

	return
}

func (m *baseCacheModel) HashGet(key, field string, data interface{}) (bool, error) {
	cache, err := m.getCache(false)

	if err != nil {
		return false, err
	}

	result, err := cache.HGet(context.Background(), key, field).Result()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(result), data)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *baseCacheModel) HashSet(key, field string, data interface{}, expire ...time.Duration) (err error) {
	cache, err := m.getCache(true)

	if err != nil {
		return
	}

	result, err := json.Marshal(data)

	if err != nil {
		return
	}

	exp := m.expire

	if len(expire) > 0 {
		exp = expire[0]
	}

	pipe := cache.Pipeline()
	pipe.HSet(context.Background(), key, field, result)

	if exp > 0 {
		pipe.Expire(context.Background(), key, exp)
	}
	_, err = pipe.Exec(context.Background())

	return
}

func (m *baseCacheModel) HashDel(key string, field ...string) (err error) {
	cache, err := m.getCache(true)

	if err != nil {
		return
	}

	if len(field) == 0 {
		return cache.Del(context.Background(), key).Err()
	}

	return cache.HDel(context.Background(), key, field...).Err()
}

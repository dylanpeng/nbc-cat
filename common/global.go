package common

import (
	"cat/common/config"
	"errors"
	"github.com/dylanpeng/golib/gorm"
	"github.com/dylanpeng/golib/http"
	"github.com/dylanpeng/golib/logger"
	"github.com/dylanpeng/golib/redis"
	oGorm "gorm.io/gorm"
)

var Logger *logger.Logger
var dbPool *gorm.Pool
var HttpServer *http.Server
var cachePool *redis.Pool

func InitLogger(config *logger.Config) (err error) {
	Logger, err = logger.NewLogger(config)
	return err
}

func InitDB(configs map[string]*gorm.Config) (err error) {
	if Logger == nil {
		return errors.New("logger uninitialized")
	}

	dbPool = gorm.NewPool(Logger)

	for k, v := range configs {
		if err := dbPool.Add(k, v); err != nil {
			return err
		}
	}

	return nil
}

func InitCache() {
	confs := config.GetConfig().Cache
	cachePool = redis.NewPool()

	for k, v := range confs {
		cachePool.Add(k, v)
	}
}

func GetDB(name string) (*oGorm.DB, error) {
	return dbPool.Get(name)
}

func InitHttpServer(router http.Router) {
	c := config.GetConfig().Http
	HttpServer = http.NewServer(c, router, Logger)
	HttpServer.Start()
}

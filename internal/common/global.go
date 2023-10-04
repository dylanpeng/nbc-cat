package common

import (
	"errors"
	"github.com/dylanpeng/golib/gorm"
	"github.com/dylanpeng/golib/logger"
	oGorm "gorm.io/gorm"
)

var Logger *logger.Logger
var dbPool *gorm.Pool

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

func GetDB(name string) (*oGorm.DB, error) {
	return dbPool.Get(name)
}

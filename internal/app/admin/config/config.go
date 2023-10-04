package config

import (
	"github.com/dylanpeng/golib/gorm"
	"github.com/dylanpeng/golib/logger"
	"github.com/jinzhu/configor"
)

var config *Config

type Config struct {
	Log *logger.Config          `toml:"log"`
	DB  map[string]*gorm.Config `toml:"db"`
}

func Init(file string) error {
	config = &Config{}

	err := configor.Load(config, file)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return config
}

package config

import (
	oConf "cat/common/config"
	"github.com/jinzhu/configor"
)

var config *Config

type Config struct {
	*oConf.Config `toml:"config" json:"config" yaml:"config"`
	Admin         *AdminConfig `toml:"admin" json:"admin" yaml:"admin"`
}

type AdminConfig struct {
	Domain string `toml:"domain" json:"domain"`
}

func Init(file string) error {
	config = &Config{
		Config: oConf.Default(),
	}

	if err := configor.Load(config, file); err != nil {
		return err
	}

	if err := config.Config.Init(); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return config
}

package config

import (
	"github.com/dylanpeng/golib/coder"
	"github.com/dylanpeng/golib/gorm"
	"github.com/dylanpeng/golib/http"
	"github.com/dylanpeng/golib/logger"
	"github.com/dylanpeng/golib/redis"
)

type App struct {
	Name      string       `toml:"name" json:"name" yaml:"name"`
	Project   string       `toml:"project" json:"project" yaml:"project"`
	Env       string       `toml:"env" json:"env" yaml:"env"`
	Debug     bool         `toml:"debug" json:"debug" yaml:"debug"`
	Secret    string       `toml:"secret" json:"secret" yaml:"secret"`
	HttpCode  string       `toml:"http_code" json:"http_code" yaml:"httpCode"`
	TcpCode   string       `toml:"tcp_code" json:"tcp_code" yaml:"tcpCode"`
	HttpCoder coder.ICoder `toml:"-" json:"-"`
	TcpCoder  coder.ICoder `toml:"-" json:"-"`
}

type Config struct {
	App   *App                     `toml:"app" json:"app" yaml:"app"`
	Http  *http.Config             `toml:"http" json:"http" yaml:"http"`
	Cache map[string]*redis.Config `toml:"cache" json:"cache" yaml:"cache"`
	DB    map[string]*gorm.Config  `toml:"db" json:"db" yaml:"db"`
	Log   *logger.Config           `toml:"log" json:"log" yaml:"log"`
}

var conf *Config

func (c *Config) Init() (err error) {
	// set coder
	if c.App.HttpCode == coder.EncodingProtobuf {
		c.App.HttpCoder = coder.ProtoCoder
	} else {
		c.App.HttpCoder = coder.JsonCoder
	}

	if c.App.TcpCode == coder.EncodingProtobuf {
		c.App.TcpCoder = coder.ProtoCoder
	} else {
		c.App.TcpCoder = coder.JsonCoder
	}

	conf = c
	return
}

func Default() *Config {
	return &Config{
		Http: http.DefaultConfig(),
		Log:  logger.DefaultConfig(),
	}
}

func GetConfig() *Config {
	return conf
}

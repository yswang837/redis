package config

import (
	"errors"
	"github.com/go-ini/ini"
	"os"
)

var (
	confRoot = os.Getenv("CONF_ROOT") + "/config/redis/"
)

type Config struct {
	Master      []string
	Slave       []string
	maxConnects int
}

func DefaultConf() *Config {
	return &Config{
		Master: []string{},
		Slave:  []string{},
	}
}

func NewConfigByFileName(name string) (*Config, error) {
	if name == "" {
		return nil, errors.New("param is empty")
	}
	f, err := ini.ShadowLoad(confRoot + name + ".ini")
	if err != nil {
		return nil, err
	}
	config, err := parseConfig(f)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func parseConfig(f *ini.File) (*Config, error) {
	conf := DefaultConf()
	conf.Master = f.Section("redis").Key("master").ValueWithShadows()
	conf.Slave = f.Section("redis").Key("slave").ValueWithShadows()
	return conf, nil
}

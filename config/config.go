package config

import (
	"bytes"
	"github.com/labstack/gommon/log"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var builtinConfig = []byte(
	`server:
  addr: 0.0.0.0:8080
mongodb:
  uri: mongodb://localhost:27017
  database_name: yaus
jwt:
  ttl: 72h
  key: somesecretekey
`)

type Config struct {
	Server  Server  `yaml:"server"`
	MongoDB MongoDB `yaml:"mongodb"`
	JWT     JWT     `yaml:"jwt"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type MongoDB struct {
	URI          string `yaml:"uri"`
	DatabaseName string `yaml:"database_name"`
}

type JWT struct {
	TTL time.Duration `yaml:"ttl"`
	Key string        `yaml:"key"`
}

func ReadConfig(filename string) *Config {
	v := viper.New()
	c := new(Config)

	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	if err := v.ReadConfig(bytes.NewReader(builtinConfig)); err != nil {
		log.Fatal("loading builtin config failed string", err)
	}

	if filename != "" {
		v.SetConfigFile(filename)
		if err := v.MergeInConfig(); err != nil {
			log.Warnf("loading config file [%s] failed: %s", filename, err)
		} else {
			log.Infof("config file [%s] loaded successfully", filename)
		}
	}

	err := v.Unmarshal(c, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
	})
	if err != nil {
		log.Fatal("failed on config unmarshal: ", err)
	}

	log.Debugf("Following configuration is loaded:\n%+v\n", c)
	return c
}

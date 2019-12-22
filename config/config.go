package config

import (
	"bytes"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var builtinConfig = []byte(`server:
  addr: 0.0.0.0:8080
mongodb:
  uri: mongodb://localhost:27017
  database_name: yaus
redis:
  address: 127.0.0.1:6379
  password: ""
  db: 0
  max_retries: 0
  min_retry_back_off: 8ms
  max_retry_back_off: 512ms
  dial_timeout: 5s
  read_timeout: 3s
  write_timeout: 3s
  pool_size: 10
  min_idle_connections: 5
  max_connection_age: 0
  pool_timeout: 4s
  idle_timeout: 5m
  idle_check_frequency: 1m
jwt:
  ttl: 72h
  key: somesecretekey
`)

type Config struct {
	Server  Server  `yaml:"server"`
	MongoDB MongoDB `yaml:"mongodb"`
	Redis   Redis   `yaml:"redis"`
	JWT     JWT     `yaml:"jwt"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type MongoDB struct {
	URI          string `yaml:"uri"`
	DatabaseName string `yaml:"database_name"`
}

type Redis struct {
	Address            string        `yaml:"address"`
	Password           string        `yaml:"password"`
	DB                 int           `yaml:"db"`
	MaxRetries         int           `yaml:"max_retries"`
	MinRetryBackOff    time.Duration `yaml:"min_retry_back_off"`
	MaxRetryBackOff    time.Duration `yaml:"max_retry_back_off"`
	DialTimeout        time.Duration `yaml:"dial_timeout"`
	ReadTimeout        time.Duration `yaml:"read_timeout"`
	WriteTimeout       time.Duration `yaml:"write_timeout"`
	PoolSize           int           `yaml:"pool_size"`
	MinIdleConnections int           `yaml:"min_idle_connections"`
	MaxConnectionAge   time.Duration `yaml:"max_connection_age"`
	PoolTimeout        time.Duration `yaml:"pool_timeout"`
	IdleTimeout        time.Duration `yaml:"idle_timeout"`
	IdleCheckFrequency time.Duration `yaml:"idle_check_frequency"`
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

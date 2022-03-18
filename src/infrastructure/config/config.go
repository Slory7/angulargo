package config

import (
	"github.com/nuveo/log"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/json"
)

type Config struct {
	AppIsDebug                 bool     `json:"appIsDebug" cfg:"appIsDebug"`
	Addr                       string   `json:"addr" cfg:"addr"`
	DBType                     string   `json:"dbType" cfg:"dbType"`
	DBConnectionString         string   `json:"dbConnectionString" cfg:"dbConnectionString"`
	DBReadOnlyConnectionString string   `json:"dbReadOnlyConnectionString" cfg:"dbReadOnlyConnectionString"`
	Redis                      RedisCfg `json:"redis" cfg:"redis"`
	CacheByMinutes             int      `json:"cacheByMinutes" cfg:"cacheByMinutes"`
	RollbackVersionID          string
}

type RedisCfg struct {
	Hosts        string `json:"hosts" cfg:"hosts"`
	Password     string `json:"password" cfg:"password"`
	MasterName   string `json:"masterName" cfg:"masterName"`
	DBNumber     int    `json:"dbNumber" cfg:"dbNumber"`
	MaxRetries   int    `json:"maxRetries" cfg:"maxRetries" cfgDefault:"0"`
	PoolSize     int    `json:"poolSize" cfg:"poolSize" cfgDefault:"1000"`
	IdleTimeout  int    `json:"idleTimeout" cfg:"idleTimeout" cfgDefault:"300"`
	MinIdleConns int    `json:"minIdleConns" cfg:"minIdleConns" cfgDefault:"0"`
}

func GetConfig[T any](environment string) T {
	if environment != "" {
		goconfig.File = "config." + environment + ".json"
	} else {
		goconfig.File = "config.json"
	}
	goconfig.DisableFlags = true
	config := new(T)
	err := goconfig.Parse(config)
	if err != nil {
		log.Fatal("config: ", err)
	}
	return *config
}

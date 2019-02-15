package test

import (
	"fmt"
	"services/infrastructure/config"
	"services/infrastructure/framework/cache"
	"services/infrastructure/framework/globals"
	"services/infrastructure/services"
	"testing"
	"time"

	_ "github.com/crgimenes/goconfig/json"
)

func initGlobal() {
	//Config
	globals.Config = config.GetConfig(globals.GetEnvironment())
	//Cache
	globals.Cache = cache.NewCache(time.Minute*120, time.Minute*5)
}

func TestTokenService(t *testing.T) {

	initGlobal()

	yhTokenService := services.NewTokenService(
		"http://api.xxx.cn",
		"/token",
		"a",
		"a",
		"a",
		"a",
		"a",
	)
	token, err := yhTokenService.GetToken()
	fmt.Printf("%v", err)
	fmt.Printf("%v", token)
}

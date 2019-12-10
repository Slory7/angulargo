package test

import (
	"fmt"
	"github.com/slory7/angulargo/src/services/infrastructure/config"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/cache"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/globals"
	"github.com/slory7/angulargo/src/services/infrastructure/services"
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

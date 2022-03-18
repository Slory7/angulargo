package test

import (
	"fmt"
	"testing"

	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/config"
	"github.com/slory7/angulargo/src/infrastructure/services"

	_ "github.com/gosidekick/goconfig/json"
)

func init() {
	//Config
	config := config.GetConfig(app.GetEnvironment(), &config.Config{}).(*config.Config)
	app.InitAppInstance(config)
	app.Instance.InitCache()
}

func TestTokenService(t *testing.T) {

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

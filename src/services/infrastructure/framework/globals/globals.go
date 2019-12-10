package globals

import (
	"os"
	"github.com/slory7/angulargo/src/services/infrastructure/config"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/cache"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/validates"

	"github.com/jwells131313/dargo/ioc"
)

var Config *config.Config

var Cache *cache.MemoryCache

var Validator *validates.Validator

var ServiceLocator ioc.ServiceLocator

func GetEnvironment() string {
	s := os.Getenv("APP_ENV")
	return s
}

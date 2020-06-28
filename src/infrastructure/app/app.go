package app

import (
	"log"
	"os"
	"time"

	"github.com/jwells131313/dargo/ioc"
	"github.com/slory7/angulargo/src/infrastructure/config"
	"github.com/slory7/angulargo/src/infrastructure/data"
	"github.com/slory7/angulargo/src/infrastructure/data/db"
	"github.com/slory7/angulargo/src/infrastructure/data/migration"
	"github.com/slory7/angulargo/src/infrastructure/framework/cache"
	"github.com/slory7/angulargo/src/infrastructure/framework/validates"
)

type App struct {
	Config *config.Config

	Cache *cache.MemoryCache

	Validator *validates.Validator

	ServiceLocator ioc.ServiceLocator

	db, dbReadOnly *db.Database
}

var Instance App

func InitAppInstance(conf *config.Config) {
	Instance = App{Config: conf}
}

func (app *App) InitDB() {
	var conf = app.Config
	app.db, app.dbReadOnly = data.NewPairDB(conf.DBType, conf.DBConnectionString, conf.DBReadOnlyConnectionString, conf.AppIsDebug)
}

func (app *App) MigrateOrRollback(init *migration.Migration, migs []*migration.Migration, rollbackVersionID string) {
	migration.MigrateOrRollback(app.db, init, migs, rollbackVersionID)
}

func (app *App) InitCache() {
	n := app.Config.CacheByMinutes
	if n <= 0 {
		n = 120
	}
	if app.Config.Redis == nil {
		app.Cache = cache.NewCache(time.Minute*time.Duration(n), time.Minute*5)
	} else {
		app.Cache = cache.NewCacheDistributed(time.Minute*time.Duration(n), time.Minute*5, *app.Config.Redis)
	}
}

func (app *App) InitValidator() {
	app.Validator = validates.NewValidator()
}

func (app *App) CacheEntity(entities ...interface{}) {
	for _, en := range entities {
		db.CacheEntity(en, *app.Config.Redis, app.db, app.dbReadOnly)
		log.Println("CacheEntity: %v", en)
	}
}

func GetEnvironment() string {
	s := os.Getenv("APP_ENV")
	return s
}

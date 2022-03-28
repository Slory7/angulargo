package main

import (
	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/config"
	"github.com/slory7/angulargo/src/services/trending/data"
	"github.com/slory7/angulargo/src/services/trending/data/migrations"
	"github.com/slory7/angulargo/src/services/trending/globals"
)

func main() {

	//Config
	globals.GlbConfig = config.GetConfig[globals.Config](app.GetEnvironment())

	app.InitAppInstance(&globals.GlbConfig.Config)

	//init db
	app.Instance.InitDB()

	//db migration
	app.Instance.MigrateOrRollback(migrations.InitMigration(), migrations.MigrationVersions, globals.GlbConfig.RollbackVersionID)

	//cache
	app.Instance.InitCache()

	//data cache
	app.Instance.CacheEntity(data.GetCacheEntities()...)

	//validator
	app.Instance.InitValidator()

	//ioc
	app.Instance.RegisterIoC(registerIoC)

	//Start rpc
	StartRpc()
}

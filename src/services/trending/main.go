package main

import (
	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/config"
	"github.com/slory7/angulargo/src/services/trending/data"
	"github.com/slory7/angulargo/src/services/trending/data/migrations"
)

func main() {

	//Config
	glbConfig = config.GetConfig(app.GetEnvironment(), &Config{}).(*Config)

	app.InitAppInstance(&glbConfig.Config)

	//init db
	app.Instance.InitDB()

	//db migration
	app.Instance.MigrateOrRollback(migrations.InitMigration(), migrations.MigrationVersions, glbConfig.RollbackVersionID)

	//cache
	app.Instance.InitCache()

	//data cache
	app.Instance.CacheEntity(data.GetCacheEntities())

	//ioc
	app.Instance.RegisterIoC(registerIoC)

	//Start rpc
	StartRpc()
}

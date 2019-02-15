package main

import (
	"services/infrastructure/appstart"
	"services/infrastructure/config"
	"services/infrastructure/data/repositories"
	"services/infrastructure/framework/cache"
	"services/infrastructure/framework/globals"
	"services/trending/data"
	"services/trending/data/migrations"
	"time"
)

func main() {

	//Config
	glbConfig = config.GetConfig(globals.GetEnvironment(), &Config{}).(*Config)

	//Cache
	globals.Cache = cache.NewCacheDistributed(time.Minute*120, time.Minute*5, glbConfig.Redis)

	//db Init
	db, dbReadOnly := appstart.InitDB(glbConfig.DBType, glbConfig.DBConnectionString, glbConfig.DBReadOnlyConnectionString, glbConfig.AppIsDebug)

	//db.Sync(new(datamodels.User))

	//db migration
	appstart.MigrationOrRollback(db, migrations.InitMigration(), migrations.MigrationVersions, glbConfig.RollbackVersionID)

	//data cache
	data.CacheEntities(db, dbReadOnly, glbConfig.Redis)

	repo := repositories.NewRepository(db)
	repoReadOnly := repositories.NewRepositoryReadOnly(dbReadOnly)

	//IoC
	RegisterIoC(repo, repoReadOnly)

	//Start rpc
	StartRpc()
}

package main

import (
	"github.com/slory7/angulargo/src/services/infrastructure/config"
	"github.com/slory7/angulargo/src/services/infrastructure/data"
	"github.com/slory7/angulargo/src/services/infrastructure/data/migration/migrations"
	"github.com/slory7/angulargo/src/services/infrastructure/data/repositories"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/cache"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/globals"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/validates"
	"time"

	"github.com/slory7/angulargo/src/services/infrastructure/appstart"

	"github.com/kataras/iris"

	"net/http"
	_ "net/http/pprof"
)

// var (
// 	rollbackVersionID = flag.String("rollback", "", "Rollback migration version id")
// )

func main() {

	//flag not compatible with goconfig
	//flag.Parse()

	//Config
	globals.Config = config.GetConfig(globals.GetEnvironment(), &config.Config{}).(*config.Config)
	conf := globals.Config

	//Cache
	globals.Cache = cache.NewCacheDistributed(time.Minute*120, time.Minute*5, conf.Redis)

	//validator
	globals.Validator = validates.NewValidator()

	//db Init
	db, dbReadOnly := appstart.InitDB(conf.DBType, conf.DBConnectionString, conf.DBReadOnlyConnectionString, conf.AppIsDebug)

	//db.Sync(new(datamodels.User))

	//db migration
	appstart.MigrationOrRollback(db, migrations.InitMigration(), migrations.MigrationVersions, conf.RollbackVersionID)

	//data cache
	data.CacheEntities(db, dbReadOnly, conf.Redis)

	repo := repositories.NewRepository(db)
	repoReadOnly := repositories.NewRepositoryReadOnly(dbReadOnly)

	//IoC
	appstart.RegisterIoC(repo, repoReadOnly)

	app := iris.New()

	//routes
	appstart.ConfigureRoutes(app)

	//curl localhost:8181/debug/pprof/trace?seconds=10 > trace.out
	//go tool trace github.com/slory7/angulargo/src/services/infrastructure/infrastructureinfrastructure.exe trace.out
	//http://www.sharelinux.com/2017/03/22/Golang%E4%B9%8Bprofiler%E5%92%8Ctrace%E5%B7%A5%E5%85%B7/
	if globals.Config.AppIsDebug {
		go func() {
			http.ListenAndServe("localhost:8181", http.DefaultServeMux)
		}()
	}

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr(conf.Addr),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

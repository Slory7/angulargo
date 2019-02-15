package appstart

import (
	"log"
	"services/infrastructure/data"
	"services/infrastructure/data/migration"
)

func InitDB(dbType string, connectionString string, readonlyConnectionString string, showSQL bool) (db *data.Database, dbReadOnly *data.Database) {
	//db
	log.Println("db type:", dbType)

	dbrw, err := data.NewDB(dbType, connectionString, showSQL)
	if err != nil {
		log.Fatal("db: ", err)
	}
	sDBReadConnectionString := readonlyConnectionString
	if len(sDBReadConnectionString) == 0 {
		sDBReadConnectionString = connectionString
	}
	dbr, err := data.NewDB(dbType, sDBReadConnectionString, showSQL)
	if err != nil {
		log.Fatal("db: ", err)
	}
	return dbrw, dbr
}

func MigrationOrRollback(db *data.Database, init *migration.Migration, migs []*migration.Migration, rollbackVersionID string) {
	//db migration
	mig := migration.New(db, &migration.Options{TableName: "app_versions", IDColumnName: "versionid"}, migs)
	mig.SetInitSchema(init)

	//rollback command
	if len(rollbackVersionID) > 0 {
		rmig := mig.GetMigration(rollbackVersionID)
		if rmig == nil {
			log.Fatal("rollback migration not exists: ", rollbackVersionID)
		}
		err := mig.RollbackMigration(rmig)
		if err != nil {
			log.Fatal("rollback migration error: ", err)
		}
	} else {
		err := mig.Migrate()
		if err != nil {
			log.Fatal("migration error: ", err)
		}
	}
}

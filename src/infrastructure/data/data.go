package data

import (
	"log"

	"github.com/slory7/angulargo/src/infrastructure/data/db"
	m "github.com/slory7/angulargo/src/infrastructure/datamodels"
)

func NewPairDB(dbType string, connectionString string, readonlyConnectionString string, showSQL bool) (dbw *db.Database, dbReadOnly *db.Database) {
	//db
	log.Println("db type:", dbType)

	dbrw, err := db.NewDB(dbType, connectionString, showSQL)
	if err != nil {
		log.Fatal("db: ", err)
	}
	sDBReadConnectionString := readonlyConnectionString
	if len(sDBReadConnectionString) == 0 {
		sDBReadConnectionString = connectionString
	}
	dbr, err := db.NewDB(dbType, sDBReadConnectionString, showSQL)
	if err != nil {
		log.Fatal("db: ", err)
	}
	return dbrw, dbr
}

func GetCacheEntities() []interface{} {
	return []interface{}{
		new(m.User),
		new(m.UserDetail),
	}
}

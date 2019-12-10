package data

import (
	"log"
	"github.com/slory7/angulargo/src/services/infrastructure/config"
	"github.com/slory7/angulargo/src/services/infrastructure/datamodels"
)

func CacheEntities(db *Database, dbReadOnly *Database, redisConf config.RedisCfg) {
	CacheEntity(new(datamodels.User), redisConf, db, dbReadOnly)
	CacheEntity(new(datamodels.UserDetail), redisConf, db, dbReadOnly)
	log.Println("CacheEntities: finished")
}

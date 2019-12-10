package data

import (
	"log"
	"github.com/slory7/angulargo/src/services/infrastructure/config"
	"github.com/slory7/angulargo/src/services/infrastructure/data"
	m "github.com/slory7/angulargo/src/services/trending/datamodels"
)

func CacheEntities(db *data.Database, dbReadOnly *data.Database, redisConf config.RedisCfg) {
	data.CacheEntity(new(m.GitRepo), redisConf, db, dbReadOnly)
	data.CacheEntity(new(m.GitRepoTrending), redisConf, db, dbReadOnly)
	log.Println("CacheEntities: finished")
}

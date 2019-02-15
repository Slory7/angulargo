package data

import (
	"log"
	"services/infrastructure/config"
	"services/infrastructure/data"
	m "services/trending/datamodels"
)

func CacheEntities(db *data.Database, dbReadOnly *data.Database, redisConf config.RedisCfg) {
	data.CacheEntity(new(m.GitRepo), redisConf, db, dbReadOnly)
	data.CacheEntity(new(m.GitRepoTrending), redisConf, db, dbReadOnly)
	log.Println("CacheEntities: finished")
}

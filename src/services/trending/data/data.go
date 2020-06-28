package data

import (
	m "github.com/slory7/angulargo/src/services/trending/datamodels"
)

func GetCacheEntities() []interface{} {
	return []interface{}{
		new(m.GitRepo),
		new(m.GitRepoTrending),
	}
}

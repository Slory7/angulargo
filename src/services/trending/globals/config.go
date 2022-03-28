package globals

import (
	"github.com/slory7/angulargo/src/infrastructure/config"

	_ "github.com/gosidekick/goconfig/json"
)

type Config struct {
	config.Config
	TrendingURL           string `json:"trendingURL" cfg:"trendingURL"`
	KeepRecentNumber      int    `json:"keepRecentNumber" cfg:"keepRecentNumber" cfgDefault:"30"`
	RequestTimeoutSeconds int    `json:"requestTimeoutSeconds" cfg:"requestTimeoutSeconds"`
}

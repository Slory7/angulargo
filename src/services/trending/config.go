package main

import (
	"github.com/slory7/angulargo/src/infrastructure/config"

	_ "github.com/gosidekick/goconfig/json"
)

type Config struct {
	config.Config
	TrendingURL           string `json:"trendingURL" cfg:"trendingURL"`
	RequestTimeoutSeconds int    `json:"requestTimeoutSeconds" cfg:"requestTimeoutSeconds"`
}

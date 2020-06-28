package main

import (
	"github.com/slory7/angulargo/src/infrastructure/config"

	_ "github.com/crgimenes/goconfig/json"
)

type Config struct {
	config.Config
	TrendingURL string `json:"trendingURL" cfg:"trendingURL"`
}

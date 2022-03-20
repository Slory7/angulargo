package main

import (
	"github.com/slory7/angulargo/src/infrastructure/framework/utils"
	"github.com/slory7/angulargo/src/infrastructure/ioc"
	gs "github.com/slory7/angulargo/src/services/trending/services/githubtrending"
)

func registerIoC(binder ioc.Binder) {
	binder.Bind(utils.GetGenericName[gs.IGithubTrendingDocService](), gs.GithubTrendingDocService{})
	binder.Bind(utils.GetGenericName[gs.IGithubTrendingService](), gs.GithubTrendingService{})
}

package main

import (
	"github.com/slory7/angulargo/src/infrastructure/framework/utils"
	"github.com/slory7/angulargo/src/services/trending/services/githubtrending"

	"github.com/jwells131313/dargo/ioc"
)

func registerIoC(binder ioc.Binder) {

	binder.Bind(utils.GetInterfaceName((*githubtrending.IGithubTrendingDocService)(nil)), githubtrending.GithubTrendingDocService{})
	binder.Bind(utils.GetInterfaceName((*githubtrending.IGithubTrendingService)(nil)), githubtrending.GithubTrendingService{})

}

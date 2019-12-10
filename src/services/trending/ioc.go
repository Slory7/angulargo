package main

import (
	"github.com/slory7/angulargo/src/services/infrastructure/data/repositories"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/globals"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/utils"
	"github.com/slory7/angulargo/src/services/trending/services/githubtrending"

	"github.com/jwells131313/dargo/ioc"
)

func RegisterIoC(
	repo repositories.IRepository,
	repoReadOnly repositories.IRepositoryReadOnly) {

	locator, err := ioc.CreateAndBind("app", func(binder ioc.Binder) error {

		binder.BindConstant(utils.GetInterfaceName((*repositories.IRepository)(nil)), repo)
		binder.BindConstant(utils.GetInterfaceName((*repositories.IRepositoryReadOnly)(nil)), repoReadOnly)

		binder.Bind(utils.GetInterfaceName((*githubtrending.IGithubTrendingDocService)(nil)), githubtrending.GithubTrendingDocService{})
		binder.Bind(utils.GetInterfaceName((*githubtrending.IGithubTrendingService)(nil)), githubtrending.GithubTrendingService{})

		//binder.BindWithCreator(LoggerServiceName, newLogger).InScope(ioc.PerLookup)

		return nil
	})
	if err == nil {
		globals.ServiceLocator = locator
	}
}

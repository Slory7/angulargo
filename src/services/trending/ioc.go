package main

import (
	"services/infrastructure/data/repositories"
	"services/infrastructure/framework/globals"
	"services/infrastructure/framework/utils"
	"services/trending/services/githubtrending"

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

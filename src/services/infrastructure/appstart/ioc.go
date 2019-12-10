package appstart

import (
	"github.com/slory7/angulargo/src/services/infrastructure/data/repositories"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/globals"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/utils"
	"github.com/slory7/angulargo/src/services/infrastructure/services/users"

	"github.com/jwells131313/dargo/ioc"
)

func RegisterIoC(
	repo repositories.IRepository,
	repoReadOnly repositories.IRepositoryReadOnly) {

	locator, err := ioc.CreateAndBind("app", func(binder ioc.Binder) error {

		binder.BindConstant(utils.GetInterfaceName((*repositories.IRepository)(nil)), repo)
		binder.BindConstant(utils.GetInterfaceName((*repositories.IRepositoryReadOnly)(nil)), repoReadOnly)

		binder.Bind(utils.GetInterfaceName((*users.IUserDetailService)(nil)), users.UserDetailService{})
		binder.Bind(utils.GetInterfaceName((*users.IUserLoginService)(nil)), users.UserLoginService{})
		binder.Bind(utils.GetInterfaceName((*users.IUserService)(nil)), users.UserService{})

		//binder.BindWithCreator(LoggerServiceName, newLogger).InScope(ioc.PerLookup)

		return nil
	})
	if err == nil {
		globals.ServiceLocator = locator
	}
}

func GetIoCInstance(interfacePointer interface{}) (interface{}, error) {
	serviceName := utils.GetInterfaceName(interfacePointer)
	service, err := globals.ServiceLocator.GetDService(serviceName)
	return service, err
}

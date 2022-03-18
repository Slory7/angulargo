package app

import (
	"log"

	"github.com/slory7/angulargo/src/infrastructure/data/repositories"
	"github.com/slory7/angulargo/src/infrastructure/framework/net/httpclient"
	"github.com/slory7/angulargo/src/infrastructure/framework/utils"

	"github.com/jwells131313/dargo/ioc"
)

func (app *App) RegisterIoC(binder func(ioc.Binder)) {

	locator, err := ioc.CreateAndBind("app", func(bi ioc.Binder) error {
		if app.db != nil {
			repo := repositories.NewRepository(app.db)
			repoReadOnly := repositories.NewRepositoryReadOnly(app.dbReadOnly)
			bi.BindConstant(utils.GetInterfaceName((*repositories.IRepository)(nil)), repo)
			bi.BindConstant(utils.GetInterfaceName((*repositories.IRepositoryReadOnly)(nil)), repoReadOnly)
		}

		bi.Bind(utils.GetInterfaceName((*httpclient.IHttpClient)(nil)), httpclient.HttpClient{})

		if binder != nil {
			binder(bi)
		}
		// binder.Bind(utils.GetInterfaceName((*users.IUserDetailService)(nil)), users.UserDetailService{})
		// binder.Bind(utils.GetInterfaceName((*users.IUserLoginService)(nil)), users.UserLoginService{})
		// binder.Bind(utils.GetInterfaceName((*users.IUserService)(nil)), users.UserService{})

		//binder.BindWithCreator(LoggerServiceName, newLogger).InScope(ioc.PerLookup)

		return nil
	})
	if err == nil {
		app.ServiceLocator = locator
	}
}

func GetIoCInstance[T any]() (T, error) {
	serviceName := utils.GetInterfaceName((*T)(nil))
	service, err := Instance.ServiceLocator.GetDService(serviceName)
	if err == nil {
		return service.(T), nil
	}
	return (any)(nil).(T), err
}

func GetIoCInstanceMust[T any]() T {
	o, err := GetIoCInstance[T]()
	if err != nil {
		log.Fatal(err)
	}
	return o
}

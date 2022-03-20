package app

import (
	"log"

	oioc "github.com/jwells131313/dargo/ioc"
	"github.com/slory7/angulargo/src/infrastructure/data/repositories"
	"github.com/slory7/angulargo/src/infrastructure/framework/net/httpclient"
	"github.com/slory7/angulargo/src/infrastructure/framework/utils"
	"github.com/slory7/angulargo/src/infrastructure/ioc"
)

//Ioc locator
var serviceLocator oioc.ServiceLocator

func (app *App) RegisterIoC(binder func(ioc.Binder)) {

	locator, err := oioc.CreateAndBind("app", func(bi oioc.Binder) error {
		if app.db != nil {
			repo := repositories.NewRepository(app.db)
			repoReadOnly := repositories.NewRepositoryReadOnly(app.dbReadOnly)
			bi.BindConstant(utils.GetInterfaceName((*repositories.IRepository)(nil)), repo)
			bi.BindConstant(utils.GetInterfaceName((*repositories.IRepositoryReadOnly)(nil)), repoReadOnly)
		}

		bi.Bind(utils.GetGenericName[httpclient.IHttpClient](), httpclient.HttpClient{})

		if binder != nil {
			mybinder := ioc.NewMyBinder()

			binder(mybinder)

			for k, v := range mybinder.BindMap {
				bi.Bind(k, v)
			}
			for k, v := range mybinder.BindWithCreatorMap {
				bi.BindWithCreator(k, func(sl oioc.ServiceLocator, d oioc.Descriptor) (interface{}, error) {
					return v()
				})
			}
			for k, v := range mybinder.BindConstantMap {
				bi.BindConstant(k, v)
			}
		}

		//binder.BindWithCreator(LoggerServiceName, newLogger).InScope(ioc.PerLookup)

		return nil
	})
	if err == nil {
		serviceLocator = locator
	}
}

func GetIoCInstance[T any]() (T, error) {
	serviceName := utils.GetInterfaceName((*T)(nil))
	service, err := serviceLocator.GetDService(serviceName)
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

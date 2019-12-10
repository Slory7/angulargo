package appstart

import (
	"github.com/slory7/angulargo/src/services/infrastructure/controllers"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/globals"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/utils"
	"github.com/slory7/angulargo/src/services/infrastructure/services/users"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func ConfigureRoutes(app *iris.Application) {
	mvc.Configure(app.Party("/user"), userPart)
}

//user controller
func userPart(app *mvc.Application) {

	serviceName := utils.GetInterfaceName((*users.IUserService)(nil))
	userService, _ := globals.ServiceLocator.GetDService(serviceName)
	app.Register(userService)

	app.Handle(new(controllers.UserController))
}

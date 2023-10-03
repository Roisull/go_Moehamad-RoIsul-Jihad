package route

import (
	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/controller"
)

func New() *echo.Echo {
	// 1. Instance echo
	e := echo.New()

	// 2. Routing
	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.RegisterController)

	return e
}

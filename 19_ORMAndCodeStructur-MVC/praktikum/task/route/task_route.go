package route

import (
	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/controller"
)

func NewRoute() *echo.Echo {
	// instance echo
	e := echo.New()

	// routing
	e.GET("/users", controller.GetUsersController)

	e.GET("/users/:id", controller.GetUserControllerById)

	e.POST("/users", controller.CreateUserController)

	e.DELETE("/users/:id", controller.DeleteUserController)

	e.PUT("/users/:id", controller.UpdateUserController)

	return e
}

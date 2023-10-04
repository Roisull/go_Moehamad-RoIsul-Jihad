package route

import (
	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/controller"
)

func NewRoute() *echo.Echo {
	// instance echo
	e := echo.New()

	// routing users
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserControllerById)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	// routing books
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookByIdController)
	e.POST("/books", controller.CreateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	e.PUT("/books/:id", controller.UpdateBookController)

	return e
}

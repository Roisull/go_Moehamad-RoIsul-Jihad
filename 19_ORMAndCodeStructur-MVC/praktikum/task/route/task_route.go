package route

import (
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/controller"
	"github.com/user/go_MoehamadRoisulJihad/20_Middleware/praktikum/constants"
	m "github.com/user/go_MoehamadRoisulJihad/20_Middleware/praktikum/middleware"
)

func NewRoute() *echo.Echo {
	// instance echo
	e := echo.New()

	// routing users
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserControllerById)
	e.POST("/users", controller.CreateUserController)
	e.POST("/login", controller.LoginUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	// routing books
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookByIdController)
	e.POST("/books", controller.CreateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	e.PUT("/books/:id", controller.UpdateBookController)

	m.LogMiddleware(e)

	// grup sendiri untuk auth basic route
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUsersController)

	// grup sendiri untuk jwt route
	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	// users
	eJwt.GET("/users", controller.GetUsersController)
	eJwt.GET("/users/:id", controller.GetUserControllerById)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)

	// books
	eJwt.GET("/books", controller.GetBooksController)
	eJwt.GET("/books/:id", controller.GetBookByIdController)
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)
	


	return e
}

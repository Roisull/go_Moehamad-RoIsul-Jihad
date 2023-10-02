package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int    `json:"id"		form:"id"`
	Age   int    `json:"age"	form:"age"`
	Email string `json:"email"	form:"email"`
	Name  string `json:"name"	form:"name"`
}

func main() {
	// 1. Instance echo
	e := echo.New()

	// 2. Routing
	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)

	e.Start(":8000")
}

// membuat controller
func UserController(e echo.Context) error {

	// get data user url param
	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("age"))

	// get data form query param
	search := e.QueryParam("search")

	sort := e.QueryParam("sort")

	user := User{
		Id:    id,
		Age:   age,
		Email: "lusiro2001@students.amikom.ac.id",
		Name:  "MOEHAMAD RO ISUL JIHAD"}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"search": search,
		"sort":   sort,
	})
}

func RegisterController(c echo.Context) error {
	// email := c.FormValue("email")
	// name := c.FormValue("name")

	user := User{}

	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		// "email" : email,
		// "name"	: name,
		"user": user,
	})
}

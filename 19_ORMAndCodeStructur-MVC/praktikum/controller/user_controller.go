package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/config"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/model"
)

// membuat controller
func GetUserController(e echo.Context) error {

	// menyiapkan tempat penyimpanan untuk database
	var users []model.User
	// pemanggilan ke dalam user sekaligus checking error
	err := config.DB.Find(&users).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully",
		"data":    users,
	})

	// // get data user url param
	// id, _ := strconv.Atoi(e.Param("id"))
	// age, _ := strconv.Atoi(e.Param("age"))

	// // get data form query param
	// search := e.QueryParam("search")

	// sort := e.QueryParam("sort")

	// user := User{
	// 	Id:    id,
	// 	Age:   age,
	// 	Email: "lusiro2001@students.amikom.ac.id",
	// 	Name:  "MOEHAMAD RO ISUL JIHAD"}

	// return e.JSON(http.StatusOK, map[string]interface{}{
	// 	"user":   user,
	// 	"search": search,
	// 	"sort":   sort,
	// })
}

func RegisterController(c echo.Context) error {
	// email := c.FormValue("email")
	// name := c.FormValue("name")

	user := model.User{}
	// binding data
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		// "email" : email,
		// "name"	: name,
		"message": "Success Input User",
		"user":    user,
	})
}

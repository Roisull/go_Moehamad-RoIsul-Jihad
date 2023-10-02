package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserByIdController(c echo.Context) error {
	// your solution here
	// mendapatkan id dari parameter id url
	id, _ := strconv.Atoi(c.Param("id"))

	dummyUsers := []User{
		{Id: 1, Name: "John Doe", Email: "john.doe@example.com", Password: "password123"},
		{Id: 2, Name: "Jane Doe", Email: "jane.doe@example.com", Password: "password456"},
		{Id: 3, Name: "Bob Smith", Email: "bob.smith@example.com", Password: "password789"},
	}
	// memasukkan data dummy ke variabel users
	users = dummyUsers

	// Mencari pengguna dengan ID yang sesuai
	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user by id",
				"user":     user,
			})
		}
	}

	// Jika tidak ada pengguna dengan ID yang sesuai, kirim respon error
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	// Mendapatkan id dari parameter id url
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "Invalid ID format",
		})
	}

	// Mencari pengguna dengan ID yang sesuai
	for i, user := range users {
		if user.Id == id {
			// Menghapus pengguna dari slice users
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user",
			})
		}
	}

	// Jika tidak ada pengguna dengan ID yang sesuai, kirim respon error
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	// Mendapatkan id dari parameter id url
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "Invalid ID format",
		})
	}

	// Mencari pengguna dengan ID yang sesuai
	for i, user := range users {
		if user.Id == id {
			// binding data baru dari request
			newUserData := new(User)
			if err := c.Bind(newUserData); err != nil {
				return err
			}

			// Memperbarui data pengguna
			users[i].Name = newUserData.Name
			users[i].Email = newUserData.Email
			users[i].Password = newUserData.Password

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user",
				"user":     users[i],
			})
		}
	}

	// Jika tidak ada pengguna dengan ID yang sesuai, kirim respon error
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserByIdController)
	e.DELETE("users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/config"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/model"
)

// get all users

func GetUsersController(c echo.Context) error {

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success get all users",
		"users":   users,
	})

}

// get user by id

func GetUserControllerById(c echo.Context) error {

	// mendapatkan id dari parameter url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}

	// Temukan pengguna dengan ID yang sesuai
	var user model.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Kirim respons JSON dengan pengguna yang ditemukan
	return c.JSON(http.StatusOK, user)

}

// create new user

func CreateUserController(c echo.Context) error {

	user := model.User{}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success create new user",

		"user": user,
	})

}

// delete user by id

func DeleteUserController(c echo.Context) error {

	// Mendapatkan id dari parameter URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}

	// Temukan pengguna dengan ID yang sesuai
	var user model.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Hapus pengguna dari database
	config.DB.Delete(&user)

	// Kirim respons JSON dengan pesan sukses
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})

}

// update user by id

func UpdateUserController(c echo.Context) error {

	// Mendapatkan id dari parameter URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}

	// Temukan pengguna dengan ID yang sesuai
	var user model.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Binding data baru dari permintaan
	newUserData := new(model.User)
	if err := c.Bind(newUserData); err != nil {
		return err
	}

	// Memperbarui data pengguna
	user.Name = newUserData.Name
	user.Email = newUserData.Email
	user.Password = newUserData.Password

	// Simpan perubahan ke database
	config.DB.Save(&user)

	// Kirim respons JSON dengan pesan sukses dan data pengguna yang diperbarui
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    user,
	})

}

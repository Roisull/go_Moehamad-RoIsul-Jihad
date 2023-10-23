package controller

import (
	"net/http"
	"strconv"
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/config"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/model"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/middleware"
)

type UserResponse struct {
	Message string
	Data	[]model.User
}

// get ID from JWT
func GetIdFromJwt(tokenString string) (string, error) {
	// Mendeklarasikan struktur Claims untuk menyimpan klaim dari token JWT
	claims := jwt.MapClaims{}

	// Mendekode token JWT
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Pastikan Anda menggunakan algoritma dan kunci yang sama yang digunakan untuk menghasilkan token
		return []byte("123jwt"), nil
	})

	if err != nil {
		return "", err
	}

	// Periksa apakah token valid dan berisi klaim yang sesuai
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(string); ok {
			return id, nil
		}
	}

	return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
}

// get all users

func GetUsersController(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")

	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is missing")
	}

	userId, err := GetIdFromJwt(token)

	if err != nil {
		return err
	}

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success get all users",
		"users":   users,
		"userId":  userId,
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

	token, err := middleware.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	userResponse := model.UserResponse{
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success create new user",
		"user": userResponse,
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

func LoginUserController(c echo.Context) error {

	// mengambil data user dulu
	user := model.User{}

	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{

			"message": "Login Failed",
			"erorr": err.Error(),
		})

	}

	// generate token yang sudah dibuat di middleware jwt
	token, err := middleware.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{

			"message": "Login Failed",
			"erorr": err.Error(),
		})
	}

	userResponse := model.UserResponse{
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}



	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "Login Successfully",
		"user": userResponse,
	})
}

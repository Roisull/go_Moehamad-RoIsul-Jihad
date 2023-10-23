package middleware

import (
	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/model"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/config"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user model.User
	err := config.DB.Where("email = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
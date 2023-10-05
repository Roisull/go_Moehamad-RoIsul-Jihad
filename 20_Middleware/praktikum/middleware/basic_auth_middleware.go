package middleware

import (
	"github.com/labstack/echo"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/config"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/model"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user model.User
	err := config.DB.Where("email = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
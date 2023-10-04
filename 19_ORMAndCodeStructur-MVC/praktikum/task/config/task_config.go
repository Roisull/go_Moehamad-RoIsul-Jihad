package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"github.com/user/go_MoehamadRoisulJihad/19_ORMAndCodeStructur-MVC/praktikum/task/model"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {

	config := model.Config{
		DB_Username: "root",
		DB_Password: "i2226915September20012020",
		DB_Port:     "3306",
		DB_Host:     "127.0.0.1",
		DB_Name:     "task_mvc",
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func InitialMigration() {

	DB.AutoMigrate(&model.User{})

}

func init() {

	InitDB()
	InitialMigration()

}

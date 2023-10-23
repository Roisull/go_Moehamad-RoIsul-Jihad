package config

import (
	"fmt"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

const DB_USER = "root"
const DB_PASS = "i2226915September20012020"
const DB_HOST = "127.0.0.1"
const DB_PORT = "8181"
const DB_NAME = "restfull"

const DB_USER_TEST = "root"
const DB_PASS_TEST = "i2226915September20012020"
const DB_HOST_TEST = "127.0.0.1"
const DB_PORT_TEST = "8181"
const DB_NAME_TEST = "restfull_test"

func InitDB(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}
func initMigrate() {
	DB.AutoMigrate(&model.User{})
}

func InitDBTest() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&model.User{})
	DB.AutoMigrate(&model.User{})
}

// func InitDB() {

// 	config := model.Config{
// 		DB_Username: "root",
// 		DB_Password: "i2226915September20012020",
// 		DB_Port:     "3306",
// 		DB_Host:     "127.0.0.1",
// 		DB_Name:     "testing_restful",
// 	}

// 	connectionString := fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

// 		config.DB_Username,
// 		config.DB_Password,
// 		config.DB_Host,
// 		config.DB_Port,
// 		config.DB_Name,
// 	)

// 	var err error

// 	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}

// }

// func InitialMigration() {

// 	DB.AutoMigrate(&model.User{})

// }

// func InitialMigrationBooks() {
// 	DB.AutoMigrate(&model.Book{})
// }

// func init() {

// 	InitDB()
// 	InitialMigration()
// 	InitialMigrationBooks()

// }

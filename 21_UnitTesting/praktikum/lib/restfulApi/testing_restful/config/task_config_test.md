package config_test

import (
	"fmt"
	"testing"

	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/config"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
)

func SetDBCredentials(username, password, port, host, dbName string) {
	DB_Username = username
	DB_Password = password
	DB_Port = port
	DB_Host = host
	DB_Name = dbName
}

func TestInitDB_Success(t *testing.T) {
	// Simpan konfigurasi dalam variabel lokal
	config := model.Config{
		DB_Username: "root",
		DB_Password: "i2226915September20012020",
		DB_Port:     "3306",
		DB_Host:     "127.0.0.1",
		DB_Name:     "testing_restful",
	}

	// Buat string koneksi
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	// Inisialisasi DB
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error initializing DB: %s", err)
	}

	// Periksa apakah DB sudah terinisialisasi
	if DB == nil {
		t.Fatal("Expected DB to be initialized, got nil")
	}

	var user model.User
	result := DB.First(&user)

	if result.Error != nil {
		t.Fatalf("Error querying database: %s", result.Error)
	}

	if result.RowsAffected != 0 {
		t.Fatal("Expected to find at least one row, got none")
	}
}

func TestInitDB_Failure(t *testing.T) {

	// Set kredensial database
	SetDBCredentials("wrong_username", "wrong_password", "3306", "127.0.0.1", "testing_restful")

	defer func() {
		recover() // Tangkap panic
	}()

	config.InitDB()

	if config.DB != nil {
		t.Fatal("Expected DB to be nil after failure, got initialized")
	}
}
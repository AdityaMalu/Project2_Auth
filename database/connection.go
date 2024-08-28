package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/AdityaMalu/auth/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:root@/auth"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}

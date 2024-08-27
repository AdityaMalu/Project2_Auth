package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() {
	dsn := "root:root@/auth"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
}

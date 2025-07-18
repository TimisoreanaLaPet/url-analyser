package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/url_analyser?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

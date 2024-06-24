package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=secretpassword dbname=money_management port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		panic(err)
	}
	return db
}

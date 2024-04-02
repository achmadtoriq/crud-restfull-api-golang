package config

import (
	"fmt"
	"go_api_module/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	sqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)

	db, err := gorm.Open(postgres.Open(sqlinfo), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Database Connected Successfully !")
	return db
}

package service

import (
	"fmt"
	"service/auth/config"
	"service/auth/entity"
	"service/auth/util/encryption"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	err      error
)

func InitDb() {
	host := config.GlobalConfig.Database.Host
	username := config.GlobalConfig.Database.Username
	port := config.GlobalConfig.Database.Port
	dbpassword := config.GlobalConfig.Database.Password
	dbname := config.GlobalConfig.Database.DBname
	if config.GlobalConfig.Security.Enabled {
		dbpassword = encryption.Decrypt(config.GlobalConfig.Database.Password)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, dbpassword, dbname, port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
	database.AutoMigrate(&entity.User{})
}

func GetDatabase() *gorm.DB {
	return database
}

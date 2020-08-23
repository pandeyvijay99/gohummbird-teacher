package db

import (
	"fmt"
	"strconv"

	"github.com/pandeyvijay99/gohummbird-teacher/config"
	"github.com/pandeyvijay99/gohummbird-teacher/models"

	"github.com/jinzhu/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	// DB, err = gorm.Open("mysql", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME")))

	if err != nil {
		panic("failed to connect database:")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.Product{}, &models.User{})
	fmt.Println("Database Migrated")
}

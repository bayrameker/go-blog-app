package db

import (
	"fmt"
	"example.com/blog-app-backend-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() (err error) {
	username := "root"
	password := "Br.$1453"
	dbName := "go-blog-db"
	// Bağlantı dizisi oluşturma
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{}, &models.Role{}, &models.Comment{})
	return err
}
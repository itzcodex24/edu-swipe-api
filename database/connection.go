package database

import (
	"fmt"
	"github.com/itzcodex24/edu-swipe-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file " + err.Error())
	}

	connection, err := gorm.Open(mysql.Open(os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/eduswipe"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database..")
	}

	if err := connection.AutoMigrate(models.User{}); err != nil {
		panic("Failed to migrate database..")
	}

	fmt.Println("Database connected..")
	DB = connection
}

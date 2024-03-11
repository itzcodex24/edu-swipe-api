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
		fmt.Errorf("error: %v", err)
	}

	connection, err := gorm.Open(mysql.Open(os.Getenv("db_user") + ":" + os.Getenv("db_password") + "@/eduswipe"))

	if err != nil {
		panic(err)
	}
	connection.AutoMigrate(models.User{})

	DB = connection

}

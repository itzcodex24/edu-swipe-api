package database

import (
	"fmt"
	"github.com/itzcodex24/edu-swipe-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open(os.Getenv("db_user"+":"+os.Getenv("db_password"+"@/eduswipe"))), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database..")
	}
	connection.AutoMigrate(models.User{})
	
	fmt.Println("Database connected..")
	DB = connection
}

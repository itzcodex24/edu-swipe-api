package database

import (
	"fmt"
	"github.com/itzcodex24/edu-swipe-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		fmt.Errorf("error: %v", err)
	}
  
	connection, err := gorm.Open(mysql.Open(os.Getenv("db_user"+":"+os.Getenv("db_password"+"@/eduswipe"))), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database..")
	}

  if err := connection.AutoMigrate(models.User{}); err != nil {
		panic(err)
	}

	fmt.Println("Database connected..")
	DB = connection
}

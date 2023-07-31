package database

import (
	"fmt"
	"os"

	"example.com/main/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	configuration.LoadEnvVariables()
	ConnectToDB()

}

func ConnectToDB() {
	var err error
	// Accessing the environment variables from .env file using os package
	// DB_URL is the key and the value is the connection string
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db
	if err != nil {
		fmt.Println(err)
		return
	}
	if DB != nil {
		fmt.Println("DB Connected Successfully")
	}
	if DB == nil {
		fmt.Println("DB not Connected Successfully")
	}
}

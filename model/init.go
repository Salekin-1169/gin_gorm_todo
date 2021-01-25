package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DBEngin *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var DB_URL = fmt.Sprintf("db path", os.Getenv("MYSQL_PASSWORD"))

	DBEngin, err = gorm.Open("mysql", DB_URL)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}

	DBEngin.AutoMigrate(&TodoModel{})
}

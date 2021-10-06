package services

import (
	"chrisevett/pointy/v1/core"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "./test.db")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	database.AutoMigrate(&core.Title{})
	DB = database
}

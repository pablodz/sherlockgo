package main

import (
	"github.com/jinzhu/gorm"
	"github.com/pablodz/sherlockgo/internal/database"
	"github.com/pablodz/sherlockgo/internal/endpoints"
	"github.com/pablodz/sherlockgo/internal/models"
	"github.com/pablodz/sherlockgo/internal/scraper"
)

func main() {
	// start database
	db, _ := database.StartDatabase()
	// start server
	go HandleServer(db)
	// download json from sherlock
	scraper.LoadData(db, "https://raw.githubusercontent.com/sherlock-project/sherlock/master/sherlock/resources/data.json")
}

func HandleServer(Db *gorm.DB) {
	// Migrate to create tables in database
	Db.AutoMigrate(&models.Sites{})
	// start Golang Echo server
	endpoints.HandleRequest(Db)
}

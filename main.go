package main

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pablodz/sherlockgo/internal/database"
	"github.com/pablodz/sherlockgo/internal/endpoints"
	"github.com/pablodz/sherlockgo/internal/models"
	"github.com/pablodz/sherlockgo/internal/scraper"
)

func main() {
	// start database
	db, _ := database.StartDatabase("sherlockgo")
	time.Sleep(time.Second * 3)
	// download json from sherlock
	scraper.LoadData(db, "https://raw.githubusercontent.com/sherlock-project/sherlock/master/sherlock/resources/data.json")
	// start scraper with username
	scraper.ScrapeThisUsername(db, "pablodz")
	// start server
	HandleServer(db)
}

func HandleServer(Db *gorm.DB) {
	// Migrate to create tables in database
	Db.AutoMigrate(&models.Sites{})
	Db.AutoMigrate(&models.Query{})
	// start Golang Echo server
	endpoints.HandleRequest(Db)
}

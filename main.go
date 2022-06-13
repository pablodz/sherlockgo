package main

import (
	"github.com/jinzhu/gorm"
	"github.com/pablodz/sherlockgo/internal/database"
	"github.com/pablodz/sherlockgo/internal/endpoints"
	"github.com/pablodz/sherlockgo/internal/models"
	"github.com/pablodz/sherlockgo/internal/scraper"
	"github.com/pablodz/sherlockgo/internal/utils"
)

func main() {
	// start database
	db, _ := database.StartDatabase(utils.MainDbName)
	// download json from sherlock
	scraper.LoadData(db, utils.TestUrl)
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

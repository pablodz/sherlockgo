package main

import (
	"github.com/pablodz/sherlockgo/internal/database"
	"github.com/pablodz/sherlockgo/internal/endpoints"
	"github.com/pablodz/sherlockgo/internal/models"
	"github.com/pablodz/sherlockgo/internal/scraper"
	"github.com/pablodz/sherlockgo/internal/utils"
	"gorm.io/gorm"
)

func main() {
	// start database
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}
	// migrate
	MigrateModels(db)
	// download json from sherlock
	scraper.LoadData(db, utils.TestUrl)
	// start Golang Echo server
	endpoints.HandleRequest()
}

func MigrateModels(db *gorm.DB) {
	// Migrate to create tables in database
	db.AutoMigrate(&models.Sites{})
	db.AutoMigrate(&models.UsernameResponse{})
}

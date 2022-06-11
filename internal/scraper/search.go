package scraper

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pablodz/sherlockgo/internal/models"
)

func ScrapeThisUsername(db *gorm.DB, username string) {
	// get all sites
	var listSites []models.Sites
	db.Find(listSites, &models.Sites{})
	// create http client
	client := &http.Client{}

	// send request for all
loopQuery:
	for index, site := range listSites {
		go doSearchOneSite(username, site, client)
		if index == 10 {
			break loopQuery
		}
	}
}

func doSearchOneSite(username string, site models.Sites, client *http.Client) {
	url := ParseFormatWithUsername(username, site)
	err := GetResponse(client, url)
	if err != nil {
		log.Println("[GetResponse][Error]", err)
	}
}

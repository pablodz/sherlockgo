package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pablodz/sherlockgo/internal/models"
	"gorm.io/gorm"
)

func ScrapeThisUsername(db *gorm.DB, username string) {
	// get all sites
	var listSites []models.Sites
	db.Find(&listSites)
	// create http client
	client := &http.Client{}

	// send request for all
loopQuery:
	for index, site := range listSites {
		go DoSearchOneSite(username, site, client)
		if index == 10 {
			break loopQuery
		}
	}
}

func DoSearchOneSite(username string, site models.Sites, client *http.Client) {
	url := ParseFormatWithUsername(username, site)
	found, statusCode, err := GetResponse(client, url, site)
	if err != nil {
		log.Println("[GetResponse][Error]", err)
	}

	if found {
		log.Println("[FOUND][YES]["+site.ErrorType+"]Searching in:", site.Sitename, "for", username, "at", url, "StatusCode:", statusCode)
	} else {
		log.Println("[FOUND][NOT]["+site.ErrorType+"]Searching in:", site.Sitename, "for", username, "at", url, "StatusCode:", statusCode)
	}

}

func DoSearchOneSiteChain(username string, site models.Sites, client *http.Client, chainResponses chan models.UsernameResponse) {
	url := ParseFormatWithUsername(username, site)
	found, statusCode, err := GetResponse(client, url, site)
	if err != nil {
		log.Println("[GetResponse][Error]", err)
	}
	log.Println("[FOUND]["+strconv.FormatBool(found)+"]["+site.ErrorType+"]Searching in:", site.Sitename, "for", username, "at", url, "StatusCode:", statusCode)
	chainResponses <- models.UsernameResponse{
		Found:            fmt.Sprint(found),
		URI:              url,
		MethodValidation: site.ErrorType,
		Username:         username,
		ResponseStatus:   statusCode,
		SiteName:         site.Sitename,
		IDSite:           site.IDSite,
		IDDownload:       "",
	}
}

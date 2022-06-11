package scraper

import (
	"log"
	"net/http"
	"strings"

	"github.com/pablodz/sherlockgo/internal/models"
)

func ParseFormatWithUsername(username string, site models.Sites) string {
	format := site.URLFormat
	return strings.Replace(format, "{}", username, -1)
}

func GetResponse(client *http.Client, errorType string, site models.Sites) (err error) {
	resp, err := client.Get(site.URLFormat)
	if err != nil {
		log.Println("[GetResponse][Error]", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("[GetResponse][Error]", errorType, ":", resp.StatusCode)
		return err
	}
	return nil
}

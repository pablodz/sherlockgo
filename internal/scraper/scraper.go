package scraper

import (
	"net/http"
	"strings"

	"github.com/pablodz/sherlockgo/internal/models"
)

func ParseFormatWithUsername(username string, site models.Sites) string {
	format := site.URLFormat
	return strings.Replace(format, "{}", username, -1)
}

func GetResponse(client *http.Client, url string) (err error) {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return err
	}
	return nil
}

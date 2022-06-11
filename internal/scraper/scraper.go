package scraper

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/pablodz/sherlockgo/internal/models"
)

func ParseFormatWithUsername(username string, site models.Sites) string {
	format := site.URLFormat
	return strings.Replace(format, "{}", username, -1)
}

func GetResponse(client *http.Client, url string, site models.Sites) (found bool, statusCode int, err error) {
	resp, err := client.Get(url)
	if err != nil {
		return false, 0, err
	}
	defer resp.Body.Close()

	if site.ErrorType == "status_code" {
		if resp.StatusCode != 200 {
			return false, resp.StatusCode, nil
		}
		return true, resp.StatusCode, nil
	} else if site.ErrorType == "message" {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		// Contains error message
		if strings.Contains(strings.ToLower(string(b)), strings.ToLower(site.ErrorMessage)) {
			return true, resp.StatusCode, nil
		} else {
			return false, resp.StatusCode, nil
		}
	} else if site.ErrorType == "response_url" {
		client = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}
		resp, err := client.Get(url)
		if err != nil {
			return false, 0, err
		}
		defer resp.Body.Close()

	} else {
		log.Println("No error type supported yet")
	}
	return false, 0, nil
}

package sitescraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GETLinkPreview() echo.HandlerFunc {
	return func(c echo.Context) error {

		link := "http://linkpreview.dev/api/v1/scrape?url=" + c.QueryParam("url")
		fmt.Println(link)
		// fetch json from link
		resp, err := http.Get(link)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		// read body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		// marshal body as json unknown struct
		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, data)
	}
}

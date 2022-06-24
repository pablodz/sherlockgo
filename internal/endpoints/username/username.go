package username

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pablodz/sherlockgo/internal/models"
	"github.com/pablodz/sherlockgo/internal/scraper"
	"gorm.io/gorm"
)

func GETByUsername(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		username := c.Param("username")
		// get all sites
		var listSites []models.Sites
		db.Find(&listSites)
		// create http client
		client := &http.Client{}
		// chain responses
		chainResponses := make(chan models.UsernameResponse)

		for _, site := range listSites {
			go scraper.DoSearchOneSiteChain(username, site, client, chainResponses)
			// if index == 10 {
			// 	break
			// }
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		enc := json.NewEncoder(c.Response())
		counter := 0
		for l := range chainResponses {
			l.IDDownload = strconv.Itoa(counter+1) + "/" + strconv.Itoa(len(listSites))
			if err := enc.Encode(l); err != nil {
				return err
			}
			c.Response().Flush()
			// time.Sleep(100 * time.Millisecond)
			counter++
			if counter == len(listSites) {
				break
			}
		}
		return nil
	}
}

func GETByUsernameAndSiteFilteredByFound(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		username := c.Param("username")
		wasFound := c.Param("found")
		var valWasFound bool
		if wasFound == "true" || wasFound == "1" {
			valWasFound = true
		} else {
			valWasFound = false
		}

		// get all sites
		var listSites []models.Sites
		db.Find(&listSites)
		// create http client
		client := &http.Client{}
		// chain responses
		chainResponses := make(chan models.UsernameResponse)

		for _, site := range listSites {
			go scraper.DoSearchOneSiteChain(username, site, client, chainResponses)
			// if index == 10 {
			// 	break
			// }
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		enc := json.NewEncoder(c.Response())
		counter := 0
		for l := range chainResponses {
			l.IDDownload = strconv.Itoa(counter+1) + "/" + strconv.Itoa(len(listSites))
			if l.Found == valWasFound {
				if err := enc.Encode(l); err != nil {
					return err
				}
				c.Response().Flush()
			}
			// time.Sleep(100 * time.Millisecond)
			counter++
			if counter == len(listSites) {
				break
			}
		}
		return nil
	}
}

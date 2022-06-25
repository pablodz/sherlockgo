package username

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/pablodz/sherlockgo/internal/database"

	"github.com/labstack/echo/v4"
	"github.com/pablodz/sherlockgo/internal/models"
	"github.com/pablodz/sherlockgo/internal/scraper"
)

// Get data by Username godoc
// @Summary Get results in streaming
// @Schemes
// @Description Get results in streaming
// @Tags Username
// @Accept json
// @Produce json
// @Success 200 {string} 200
// @Param username   path string example "Username"
// @Router /username/{username} [get]
func GETByUsernameStreaming() echo.HandlerFunc {
	return func(c echo.Context) error {

		username := c.Param("username")
		// get all sites
		var listSites []models.Sites
		db, err := database.GetDB()

		if err != nil {

			return err
		}

		log.Println("PRIOR DB CALL")
		db.Find(&listSites)
		log.Println("post DB CALL")
		log.Println(listSites)
		// create http client
		client := &http.Client{}
		// chain responses
		chainResponses := make(chan models.UsernameResponse)

		for _, site := range listSites {
			go scraper.DoSearchOneSiteChain(username, site, client, chainResponses)

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

// Get data by Username godoc
// @Summary Get results in streaming
// @Schemes
// @Description Get results in streaming
// @Tags Username
// @Accept json
// @Produce json
// @Success 200 {string} 200
// @Param username		path string example "Username"
// @Param found   		path bool true	"Found"
// @Router /username/{username}/found/{found} [get]
func GETByUsernameAndSiteFilteredByFoundStreaming() echo.HandlerFunc {
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

		db, err := database.GetDB()

		if err != nil {

			return err
		}

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
			if l.Found == valWasFound {
				l.IDDownload = strconv.Itoa(counter+1) + "/" + strconv.Itoa(len(listSites))
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

package username

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
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
// @Param username   	path	string	example	"Username"
// @Param limit   		path	int		10		"Limit"
// @Param format   		path	string	csv		"OutputFormat"
// @Router /username/{username} [get]
func GETByUsernameStreaming() echo.HandlerFunc {
	return func(c echo.Context) error {

		username := c.Param("username")
		format := c.QueryParam("format")
		// found := c.QueryParam("found")
		limitInt, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			return returnError(c, err)
		}
		// print all query params
		for i, v := range c.QueryParams() {
			log.Println("PARAM", i, v)
		}

		// set all not specified params
		if format == "" {
			format = "json"
		}

		// get all sites
		var listSites []models.Sites
		db, err := database.GetDB()
		if err != nil {
			return err
		}
		db.Find(&listSites)
		client := &http.Client{}
		// chain responses
		chainResponses := make(chan models.UsernameResponse)

		for idx, site := range listSites {
			go scraper.DoSearchOneSiteChain(username, site, client, chainResponses)
			if limitInt == idx {
				break
			}
		}
		// if not specified, get all
		if limitInt == 0 {
			limitInt = len(listSites) - 1
		}

		switch format {
		case "json":
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		case "csv":
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		case "xml":
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
		}
		c.Response().WriteHeader(http.StatusOK)
		counter := 1
		switch format {
		case "json":
			enc := json.NewEncoder(c.Response())
			for oneResp := range chainResponses {
				oneResp.IDDownload = strconv.Itoa(counter+1) + "/" + strconv.Itoa(limitInt)
				if err := enc.Encode(oneResp); err != nil {
					return returnError(c, err)
				}
				c.Response().Flush()
				counter++
				if counter == limitInt {
					break
				}
			}
		case "csv":
			enc := csv.NewWriter(c.Response())
			for oneResp := range chainResponses {
				oneResp.IDDownload = strconv.Itoa(counter+1) + "/" + strconv.Itoa(limitInt)
				if err := enc.Write([]string{oneResp.Found,
					oneResp.URI,
					oneResp.MethodValidation,
					oneResp.Username,
					strconv.Itoa(oneResp.ResponseStatus),
					oneResp.SiteName,
					strconv.Itoa(oneResp.IDSite),
					oneResp.IDDownload}); err != nil {
					return returnError(c, err)
				}
				enc.Flush()
				counter++
				if counter == limitInt {
					break
				}
			}
		case "xml":
			enc := xml.NewEncoder(c.Response())
			// write xml header
			if _, err := c.Response().Write([]byte(xml.Header)); err != nil {
				return returnError(c, err)
			}
			// Object General
			if _, err := c.Response().Write([]byte("<Data>")); err != nil {
				return returnError(c, err)
			}
			for oneResp := range chainResponses {
				oneResp.IDDownload = strconv.Itoa(counter+1) + "/" + strconv.Itoa(limitInt)
				if err := enc.Encode(oneResp); err != nil {
					return returnError(c, err)
				}
				c.Response().Flush()
				counter++
				if counter == limitInt {
					break
				}
			}
			if _, err := c.Response().Write([]byte("</Data>")); err != nil {
				return returnError(c, err)
			}
		}
		return nil
	}
}

func returnError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, err.Error())
}

// // Get data by Username godoc
// // @Summary Get results in streaming
// // @Schemes
// // @Description Get results in streaming
// // @Tags Username
// // @Accept json
// // @Produce json
// // @Success 200 {string} 200
// // @Param username		path string example "Username"
// // @Param found   		path bool true	"Found"
// // @Router /username/{username}/found/{found} [get]
// func GETByUsernameAndSiteFilteredByFoundStreaming() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		username := c.Param("username")
// 		wasFound := c.QueryParam("found")
// 		var valWasFound bool
// 		if wasFound == "true" || wasFound == "1" {
// 			valWasFound = true
// 		} else {
// 			valWasFound = false
// 		}

// 		// get all sites
// 		var listSites []models.Sites

// 		db, err := database.GetDB()

// 		if err != nil {

// 			return err
// 		}

// 		db.Find(&listSites)
// 		// create http client
// 		client := &http.Client{}
// 		// chain responses
// 		chainResponses := make(chan models.UsernameResponse)

// 		for _, site := range listSites {
// 			go scraper.DoSearchOneSiteChain(username, site, client, chainResponses)
// 			// if index == 10 {
// 			// 	break
// 			// }
// 		}

// 		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		c.Response().WriteHeader(http.StatusOK)

// 		enc := json.NewEncoder(c.Response())
// 		counter := 0
// 		for l := range chainResponses {
// 			if l.Found == valWasFound {
// 				l.IDDownload = strconv.Itoa(counter+1) + "/" + strconv.Itoa(len(listSites))
// 				if err := enc.Encode(l); err != nil {
// 					return err
// 				}
// 				c.Response().Flush()
// 			}
// 			// time.Sleep(100 * time.Millisecond)
// 			counter++
// 			if counter == len(listSites) {
// 				break
// 			}
// 		}
// 		return nil
// 	}
// }

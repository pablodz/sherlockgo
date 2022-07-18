package username

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strconv"

	"github.com/pablodz/sherlockgo/internal/database"
	"github.com/pablodz/sherlockgo/internal/endpoints/responses"

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
			limitInt = 10
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
					return responses.ReturnError(c, err)
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
					return responses.ReturnError(c, err)
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
				return responses.ReturnError(c, err)
			}
			// Object General
			if _, err := c.Response().Write([]byte("<Data>")); err != nil {
				return responses.ReturnError(c, err)
			}
			for oneResp := range chainResponses {
				oneResp.IDDownload = strconv.Itoa(counter+1) + "/" + strconv.Itoa(limitInt)
				if err := enc.Encode(oneResp); err != nil {
					return responses.ReturnError(c, err)
				}
				c.Response().Flush()
				counter++
				if counter == limitInt {
					break
				}
			}
			if _, err := c.Response().Write([]byte("</Data>")); err != nil {
				return responses.ReturnError(c, err)
			}
		}
		return nil
	}
}

func GETByMultipleUsernames() echo.HandlerFunc{
	
}
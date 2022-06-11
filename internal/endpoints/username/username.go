package username

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pablodz/sherlockgo/internal/scraper"
)

func GETByUsername(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		username := c.Param("username")
		scraper.ScrapeThisUsername(db, username)

		return c.JSON(http.StatusOK, "ok")

	}
}

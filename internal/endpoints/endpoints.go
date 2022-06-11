package endpoints

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pablodz/sherlockgo/internal/endpoints/sites"
)

func HandleRequest(db *gorm.DB) {
	e := echo.New()

	/* Add here the middlewares */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/* Add here the routes or endpoints */
	e.GET("/", GETsimpleResponse())
	e.GET("/sites", sites.GETListSites(db))
	e.Logger.Fatal(e.Start("0.0.0.0:6969"))
}

func GETsimpleResponse() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, I'm Sherlock, but faster!")
	}
}

package endpoints

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pablodz/sherlockgo/internal/endpoints/sites"
	"github.com/pablodz/sherlockgo/internal/endpoints/username"
)

func HandleRequest() {
	e := echo.New()

	/* Add here the middlewares */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/* Add here the routes or endpoints */
	e.GET("/", GETsimpleResponse())
	e.GET("/sites", sites.GETListSites())
	e.GET("/username/:username", username.GETByUsernameStreaming())
	e.GET("/username/:username/found/:found", username.GETByUsernameAndSiteFilteredByFoundStreaming())
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}
	e.Logger.Fatal(e.Start("0.0.0.0:" + port)) // :6969
}

func GETsimpleResponse() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, I'm Sherlock, but faster!")
	}
}

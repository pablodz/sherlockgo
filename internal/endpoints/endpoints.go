package endpoints

import (
	"net/http"
	"os"

	// "time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pablodz/sherlockgo/docs"
	"github.com/pablodz/sherlockgo/internal/endpoints/sites"
	"github.com/pablodz/sherlockgo/internal/endpoints/username"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func HandleRequest() {
	e := echo.New()

	/* Docs */
	docs.SwaggerInfo.Title = "SherlockGo API"
	docs.SwaggerInfo.Description = "This is a simple API to search for usernames in websites"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v2"
	docs.SwaggerInfo.Host = "sherlockgo.herokuapp.com"
	docs.SwaggerInfo.Schemes = []string{"https"} //"http" not supported by heroku

	/* Add here the middlewares */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Middleware
	// e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	// 	Timeout: 60 * time.Second,
	// }))

	/* Add here the routes or endpoints */
	e.GET("/api/v2/", GETsimpleResponse())
	e.GET("/api/v2/sites", sites.GETListSites())
	e.GET("/api/v2/username/:username", username.GETByUsernameStreaming())
	e.GET("/api/v2/username/:username/found/:found", username.GETByUsernameAndSiteFilteredByFoundStreaming())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}
	e.Logger.Fatal(e.Start("0.0.0.0:" + port)) // :6969
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description Do ping
// @Tags Status
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router / [get]
func GETsimpleResponse() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, I'm Sherlock, but faster!")
	}
}

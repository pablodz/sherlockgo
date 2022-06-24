package sites

import (
	"net/http"

	"github.com/pablodz/sherlockgo/internal/database"

	"github.com/labstack/echo/v4"
	"github.com/pablodz/sherlockgo/internal/models"
)

func GETListSites() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paths []models.Sites
		database.DB.Find(&paths)
		return c.JSON(http.StatusOK, paths)
	}
}

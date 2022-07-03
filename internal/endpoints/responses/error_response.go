package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReturnError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, err.Error())
}

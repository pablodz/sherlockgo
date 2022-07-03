package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReturnError(c echo.Context, err error) error {
	// get ip address
	// ip := c.RealIP()
	// if ip == "" {
	// 	ip = c.Request().RemoteAddr
	// }
	return c.JSON(http.StatusInternalServerError, err.Error())
}

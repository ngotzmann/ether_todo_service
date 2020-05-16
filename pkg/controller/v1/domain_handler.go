package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Unrestricted(c echo.Context) error {
	return c.JSON(http.StatusOK, "Unrestricted")
}

func Restricted(c echo.Context) error {
	return c.JSON(http.StatusOK, "Restricted")
}

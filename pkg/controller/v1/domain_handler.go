package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)


func Unrestricted(c echo.Context) error {
	return c.JSON(http.StatusOK, "Unrestricted")
}

func Restricted(c echo.Context) error {
	return c.JSON(http.StatusOK, "Unrestricted")
}

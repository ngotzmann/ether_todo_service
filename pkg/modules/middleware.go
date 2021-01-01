package modules

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomEchoHTTPErrorHandler(err error, c echo.Context) {
	if err != nil {
		if err := c.JSON(http.StatusBadRequest, err.Error()); err != nil {
			c.Logger().Error(err)
		}
	} else {
		if err := c.JSON(http.StatusInternalServerError, err.Error()); err != nil {
			c.Logger().Error(err)
		}
	}
	c.Logger().Error(err)
}

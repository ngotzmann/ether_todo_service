package v1

import (
	"ether_todo/pkg/injector"
	"ether_todo/pkg/todo"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var uc = injector.TodoUsecase()

func Endpoints(e *echo.Echo) *echo.Echo {
	e.GET("/todo/list/:name", FindListByName)
	e.POST("/todo/list", SaveList)
	return e
}

func FindListByName(c echo.Context) error {
	name := c.Param("name")

	l, err := uc.FindListByName(name)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, l)
	}
}

func SaveList(c echo.Context) error {
	l := &todo.List{}
	if err := c.Bind(l); err != nil {
		return err
	}

	fmt.Println(l)

	l, err := uc.SaveList(l)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, l)
	}
}

func CleanOutatedLists(c echo.Context) error {
	uc.CleanOutatedLists()
	return c.JSON(http.StatusOK, "frutti")
}
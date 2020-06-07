package v1

import (
	"ether_todo/pkg/controller/persistence"
	"ether_todo/pkg/todo"
	"github.com/ngotzmann/gorror"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FindListByName(c echo.Context) error {
	name := c.Param("name")
	repo := persistence.NewTodoListRepo()
	uc := todo.NewUsecase(repo, todo.NewService(repo))
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
		return gorror.CreateError(gorror.InternalServerError, err.Error())
	}

	repo := persistence.NewTodoListRepo()
	uc := todo.NewUsecase(repo, todo.NewService(repo))
	l, err := uc.SaveList(l)

	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, l)
	}
}
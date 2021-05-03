package v1

import (
	"ether_todo/pkg/injector"
	"ether_todo/pkg/todo"
	"ether_todo/pkg/todo/controller/v1/ws"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var uc = injector.TodoUsecase()

func Endpoints(e *echo.Echo) *echo.Echo {
	e.GET("/todo/list/:name", FindListByName)
	e.POST("/todo/list", SaveList)
	e.GET("/ws/:clientName", WebsocketHandler)
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
	l, err := uc.SaveList(l)

	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, l)
	}
}

var h = &ws.Hub{}
func WebsocketHandler(c echo.Context) error {
	upgrader := getUpgrader()
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	cltN := c.Param("clientName")
	clt := ws.Client{
		Id: cltN,
		Connection: conn,
	}

	h.AddClient(clt)
	//TODO: get error over channel
	go h.Run(conn, clt)
	return nil
}

func getUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {return true}, //TODO: Fix this dont return true all time!
		EnableCompression: true,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}
package v1

import (
	"ether_todo/pkg/injector"
	"ether_todo/pkg/todo"
	"ether_todo/pkg/todo/controller/v1/ws"
	"ether_todo/pkg/todo/controller/v1/ws2"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var uc = injector.TodoUsecase()

func Endpoints(e *echo.Echo) *echo.Echo {
	e.GET("/todo/list/:name", FindListByName)
	e.POST("/todo/list", SaveList)
	e.GET("/ws/:subId&:topic:", GlueGenericWebsocketHandler)
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


var h = &ws2.Hub{}
//TODO: Rename
func GlueGenericWebsocketHandler(c echo.Context) error {
	u := getUpgrader()
	conn, err := u.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	//clt := getClient(c, conn)
	//h.AddClient(clt)
	//go h.Handler(conn, clt)

	sub := getSubscription(c, conn)
	//TODO: possible to move in h.Handler?
	h.AddSubscription(sub)

	//TODO: handle error
	go h.Handler(conn, sub)
	return nil
}

func getSubscription(c echo.Context, conn *websocket.Conn) ws2.Subscription {
	id := c.Param("subId")
	topic := c.Param("topic")
	sub := ws2.Subscription{
		Id: id,
		Topic: topic,
		Connection: conn,
	}
	return sub
}

func getUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {return true}, //TODO: Fix this dont return true all time!
		EnableCompression: true,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

func getClient(c echo.Context, conn *websocket.Conn) ws.Client {
	cltN := c.Param("clientName")
	clt := ws.Client{
		Id: cltN,
		Connection: conn,
	}
	return clt
}
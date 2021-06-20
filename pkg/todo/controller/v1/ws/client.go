package ws

import "github.com/gorilla/websocket"

type Client struct {
	Id         string
	Connection *websocket.Conn
}

func (clt *Client) send(messageType int, message string) error {
	return clt.Connection.WriteMessage(messageType, []byte(message))
}

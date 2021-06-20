package ws2

import "github.com/gorilla/websocket"

type Subscription struct {
	Id         string
	Topic  	   string
	Connection *websocket.Conn
}

func (s *Subscription) send(messageType int, message string) error {
	return s.Connection.WriteMessage(messageType, []byte(message))
}

func (s *Subscription) validate() error {
	if s.Id == "" || s.Topic == "" || s.Connection == nil {
		//TODO: return Error
		return nil
	}
	return nil
}

func (s *Subscription) containsMe(subs []Subscription) (bool, int) {
	for i, sub := range subs {
		if s.Id == sub.Id && sub.Topic == s.Topic {
			return true, i
		}
	}
	return false, 0
}
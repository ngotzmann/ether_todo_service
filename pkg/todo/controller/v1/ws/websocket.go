package ws

import (
	"encoding/json"
	"ether_todo/pkg/todo"
	"github.com/gorilla/websocket"
	"log"
)

const (
	PUBLISH     = "publish"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

type Hub struct {
	Clients       []Client
	Subscriptions []Subscription
}

type Client struct {
	Id         string
	Connection *websocket.Conn
}

type Subscription struct {
	Topic  string
	Client *Client
}

type Request struct {
	Action  string `json:"action"`
	Topic   string `json:"topic"`
	Message todo.Task
}

//TODO: write error in channel
func (h *Hub) Run(conn *websocket.Conn, clt Client) error {
	for {
		mType, p, err := conn.ReadMessage()
		if err != nil {
			h.RemoveClient(clt)
			return err
		}
		err = h.HandleReceiveMessage(clt, mType, p)
		if err != nil {
			return err
		}
	}
}

func (h *Hub) HandleReceiveMessage(clt Client, messageType int, payload []byte) error {
	m := Request{}
	err := json.Unmarshal(payload, &m)
	if err != nil {
		return err
	}

	switch m.Action {

	case SUBSCRIBE:
		h.Subscribe(clt, m.Topic)
		break
	case PUBLISH:
		err := h.Publish(m.Topic, messageType, m.Message)
		if err != nil {
			return err
		}
		break
	case UNSUBSCRIBE:
		h.Unsubscribe(clt, m.Topic)
		break
	default:
		break
	}
	return nil
}

func (h *Hub) Publish(topic string, messageType int, message string) error {
	subs := h.GetSubscriptionsByTopic(topic)
	for _, sub := range subs {
		err := sub.Client.Send(messageType, message)
		if err != nil {
			return err
		}
	}
	return nil
}
func (clt *Client) Send(messageType int, message string) error {
	return clt.Connection.WriteMessage(messageType, []byte(message))
}

func (h *Hub) Subscribe(clt Client, topic string) {
	cltSubs := h.GetSubscriptionsOfClient(topic, clt)
	if len(cltSubs) > 0 {
		// client is subscribed this topic before
		return
	}
	newSub := Subscription{
		Topic:  topic,
		Client: &clt,
	}
	h.Subscriptions = append(h.Subscriptions, newSub)
	log.Println(clt.Id + " subscribed to topic: " + topic)
}

func (h *Hub) GetSubscriptionsByTopic(t string) []Subscription {
	var rslt []Subscription
	for _, sub := range h.Subscriptions {
		if sub.Topic == t {
			rslt = append(rslt, sub)
		}
	}
	return rslt
}




func (h *Hub) GetSubscriptionsOfClient(topic string, clt Client) []Subscription {
	var rslt []Subscription
	for _, sub := range h.Subscriptions {
		if sub.Client.Id == clt.Id && sub.Topic == topic {
			rslt = append(rslt, sub)
		}
	}
	return rslt
}

func (h *Hub) AddClient(clt Client) {
	h.Clients = append(h.Clients, clt)
}

func (h *Hub) RemoveClient(clt Client) {
	h.removeClientSubscriptions(clt)
	h.removeClientFromList(clt)
	log.Println("Connected clients and subscriptions ", len(h.Clients), len(h.Subscriptions))
}
func (h *Hub) removeClientFromList(clt Client) {
	for index, c := range h.Clients {
		if c.Id == clt.Id {
			h.Clients = append(h.Clients[:index], h.Clients[index+1:]...)
		}
	}
}
func (h *Hub) removeClientSubscriptions(clt Client) {
	for i, sub := range h.Subscriptions {
		if clt.Id == sub.Client.Id {
			h.Subscriptions = append(h.Subscriptions[:i], h.Subscriptions[i+1:]...)
		}
	}
}

func (h *Hub) Unsubscribe(clt Client, topic string) {
	cltSubs := h.GetSubscriptionsOfClient(topic, clt)
	for i, sub := range cltSubs {
		if sub.Client.Id == clt.Id && sub.Topic == topic {
			h.Subscriptions = append(h.Subscriptions[:i], h.Subscriptions[i+1:]...)
		}
	}
	log.Println(clt.Id + " was unsubscribed of topic: " + topic)
}


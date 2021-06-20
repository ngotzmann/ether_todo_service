package ws

import (
	"encoding/json"
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

type Subscription struct {
	Topic  string //listName
	Client *Client
}

type Request struct {
	Action  string `json:"action"`
	Topic   string `json:"topic"`
	ObjectJson string `json:"objectJson"`
}

//TODO: write error in channel
//TODO: Rename method
func (h *Hub) Handler(conn *websocket.Conn, clt Client) error {
	for {
		mType, p, err := conn.ReadMessage()
		if err != nil {
			h.removeClient(clt)
			return err
		}
		err = h.handleReceiveMessage(clt, mType, p)
		if err != nil {
			return err
		}
	}
}

func (h *Hub) AddClient(clt Client) {
	h.Clients = append(h.Clients, clt)
}

func (h *Hub) handleReceiveMessage(clt Client, messageType int, payload []byte) error {
	r := Request{}
	err := json.Unmarshal(payload, &r)
	if err != nil {
		return err
	}
	switch r.Action {
	case SUBSCRIBE:
		h.subscribe(clt, r.Topic)
		break
	case PUBLISH:
		err := h.publish(r.Topic, messageType, r.ObjectJson)
		if err != nil {
			return err
		}
		break
	case UNSUBSCRIBE:
		h.unsubscribe(clt, r.Topic)
		break
	default:
		break
	}
	return nil
}

func (h *Hub) publish(topic string, messageType int, message string) error {
	subs := h.getSubscriptionsByTopic(topic)
	for _, sub := range subs {
		err := sub.Client.send(messageType, message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Hub) subscribe(clt Client, topic string) {
	cltSubs := h.getSubscriptionsOfClient(topic, clt)
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

func (h *Hub) getSubscriptionsByTopic(t string) []Subscription {
	var rslt []Subscription
	for _, sub := range h.Subscriptions {
		if sub.Topic == t {
			rslt = append(rslt, sub)
		}
	}
	return rslt
}

//TODO move this method to client object
func (h *Hub) getSubscriptionsOfClient(topic string, clt Client) []Subscription {
	var rslt []Subscription
	for _, sub := range h.Subscriptions {
		if sub.Client.Id == clt.Id && sub.Topic == topic {
			rslt = append(rslt, sub)
		}
	}
	return rslt
}

func (h *Hub) removeClient(clt Client) {
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

func (h *Hub) unsubscribe(clt Client, topic string) {
	cltSubs := h.getSubscriptionsOfClient(topic, clt)
	for i, sub := range cltSubs {
		if sub.Client.Id == clt.Id && sub.Topic == topic {
			h.Subscriptions = append(h.Subscriptions[:i], h.Subscriptions[i+1:]...)
		}
	}
	log.Println(clt.Id + " was unsubscribed of topic: " + topic)
}

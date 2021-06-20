package ws2

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
	Subscriptions []Subscription
}

type Request struct {
	Action     string `json:"action"`
	Topic      string `json:"topic"`
	ObjectJson string `json:"objectJson"`
}

// TODO: write error in channel
//TODO: Rename method
//Every request(subscription) creates a new goroutine
func (h *Hub) Handler(conn *websocket.Conn, sub Subscription) error {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			h.removeSubscription(sub)
			return err
		}
		err = h.handleReceiveMessage(sub, messageType, p)
		if err != nil {
			return err
		}
	}
}

func (h *Hub) handleReceiveMessage(sub Subscription, messageType int, payload []byte) error {
	r := Request{}
	err := json.Unmarshal(payload, &r)
	if err != nil {
		return err
	}

	switch r.Action {
		case SUBSCRIBE:
			err := h.subscribe(sub)
			if err != nil {
				return err
			}
			break
		case PUBLISH:
			err := h.publish(messageType, r.ObjectJson)
			if err != nil {
				return err
			}
			break
		case UNSUBSCRIBE:
			err := h.unsubscribe(sub)
			if err != nil {
				return err
			}
			break
		default:
			break
	}
	return nil
}

func (h *Hub) subscribe(givenSub Subscription) error {
	err := givenSub.validate()
	if err != nil {
		return err
	}

	isContained, _ := givenSub.containsMe(h.Subscriptions)
	if isContained {
		//TODO: return Error
	}
	h.Subscriptions = append(h.Subscriptions, givenSub)
	log.Println(givenSub.Id + " subscribed to topic: " + givenSub.Topic)
	return nil
}

func (h *Hub) publish(messageType int, message string) error {
	//1. Build up or use an existince channel
	//2. Write request object in this channel
	//3. Listen in the sub.send method to messages which come from channel
	for _, sub := range h.Subscriptions {
		err := sub.send(messageType, message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Hub) removeSubscription(givenSub Subscription) {
	for i, sub := range h.Subscriptions {
		if givenSub.Id == sub.Id {
			h.Subscriptions = append(h.Subscriptions[:i], h.Subscriptions[i+1:]...)
		}
		log.Println("Connected subscriptions ", len(h.Subscriptions))
	}
}

func (h *Hub) unsubscribe(givenSub Subscription) error {
	err := givenSub.validate()
	if err != nil {
		return err
	}

	isContained, index := givenSub.containsMe(h.Subscriptions)
	if isContained {
		h.Subscriptions = append(h.Subscriptions[:index], h.Subscriptions[index+1:]...)
	}
	log.Println("Subscription was removed: " + givenSub.Id)
	return nil
}


func (h *Hub) AddSubscription(sub Subscription) {
	h.Subscriptions = append(h.Subscriptions, sub)
}
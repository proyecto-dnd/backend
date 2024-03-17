package ws

import (
	"encoding/json"
	"time"
)

const (
	EventSendMessage = "SEND_MESSAGE"
	EventNewMessage = "NEW_MESSAGE"
	EventChangeRoom = "CHANGE_ROOM"
)

type Event struct {
	Type    string `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent"`
}

func SendMessageHandler(event Event, c *Client) error {
	var chatEvent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatEvent); err != nil {
		return err
	}
	
	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatEvent.Message
	broadMessage.From = chatEvent.From

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return err
	}

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewMessage
	for client := range c.manager.clients {
		if client.chatroom == c.chatroom {
			client.egress <- outgoingEvent
		}
	}

	return nil
}

type ChangeRoomEvent struct {
	ChatRoomId int `json:"chat_room_id"`
}

func ChatRoomHandler(event Event, c *Client) error {
	var changeRoomEvent ChangeRoomEvent
	if err := json.Unmarshal(event.Payload, &changeRoomEvent); err != nil {
		return err
	}
	
	c.chatroom = changeRoomEvent.ChatRoomId

	return nil
}
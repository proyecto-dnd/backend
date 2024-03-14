package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[*Client]bool

type Client struct {
	conn     *websocket.Conn
	manager  *Manager
	egress   chan Event
	chatroom int
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		conn:     conn,
		manager:  manager,
		egress:   make(chan Event),
	}
}

func (c *Client) readMessage() {
	defer func ()  {
		c.manager.removeClient(c)
	}()
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	c.conn.SetReadLimit(512)

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}
		
	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println(err)
			break
		}
		
		if err := c.manager.routeEvent(request, c); err != nil {
			log.Println(err)
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	log.Println("pong")
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}

func (c *Client) writeMessages() {
	ticker := time.NewTicker(pingInterval)
	defer func ()  {
		ticker.Stop()
		c.manager.removeClient(c)
	}()

	for {
		select {
		case message, ok := <- c.egress:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}

				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}

			log.Println("message sent")
		case <- ticker.C:
			log.Println("ping")
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("write message", err)
				return
			}
		}
	}
}
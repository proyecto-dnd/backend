package ws

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				if message.SessionID == client.sessionId {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

func (h *Hub) ServeWs(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	
	sessionId, err := strconv.Atoi(ctx.Param("session_id"))
	if err != nil {
		log.Println(err)
		return
	}
	
	client := NewClient(h, conn, sessionId)
	h.register <- client

	log.Println("Nuevo cliente conectado: ", client.sessionId)

	go client.readPump()
	go client.writePump()
}
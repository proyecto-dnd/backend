package ws

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		CheckOrigin: checkOrigin,
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

func checkOrigin(r *http.Request) bool {
	return true
}

type Manager struct {
	clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
}

func NewManager() *Manager {
	m := &Manager{
		clients: make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()
	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessageHandler
	m.handlers[EventChangeRoom] = ChatRoomHandler
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) ServeWS(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.conn.Close()
		delete(m.clients, client)
	}
}
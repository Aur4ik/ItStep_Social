package handler

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	chatID string
	userID int
}

type BroadcastMessage struct {
	chatID  string
	payload interface{}
}

var (
	clients   = make(map[*Client]bool)
	clientsMu sync.Mutex
	broadcast = make(chan BroadcastMessage, 256)
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(c *gin.Context) {
	// Маршрут защищён AuthMiddleWare — user_id гарантированно есть
	userID := c.GetInt("user_id")
	chatID := c.Query("chat_id")

	if chatID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chat_id required"})
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	client := &Client{conn: ws, chatID: chatID, userID: userID}

	clientsMu.Lock()
	clients[client] = true
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, client)
		clientsMu.Unlock()
	}()

	for {
		var msg interface{}
		if err := ws.ReadJSON(&msg); err != nil {
			break
		}
		broadcast <- BroadcastMessage{chatID: chatID, payload: msg}
	}
}

func HandleMessages() {
	for msg := range broadcast {
		clientsMu.Lock()
		for client := range clients {
			if client.chatID != msg.chatID {
				continue
			}
			if err := client.conn.WriteJSON(msg.payload); err != nil {
				client.conn.Close()
				delete(clients, client)
			}
		}
		clientsMu.Unlock()
	}
}

package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handler struct {
	Hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		Hub: hub,
	}
}

func (h *Handler) HandleConnections(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		Conn: conn,
		Hub:  h.Hub,
		Send: make(chan []byte),
	}

	h.Hub.Register <- client

	defer func() {
		h.Hub.Unregister <- client
		conn.Close()
	}()

	go client.WritePump() // 👈 Add this

	client.ReadPump()
}
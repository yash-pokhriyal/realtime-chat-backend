package websocket

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
	Hub  *Hub
	Send chan []byte
}

func (c *Client) ReadPump() {

	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {

		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		c.Hub.Broadcast <- message
	}
}

func (c *Client) WritePump() {

	defer c.Conn.Close()

	for {
		message, ok := <-c.Send

		if !ok {
			return
		}

		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}
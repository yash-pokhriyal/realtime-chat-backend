// type Hub struct

// Hub ek manager hai.

// Jaise school me principal saare students ko manage karta hai.

// Waise hi Hub saare connected clients ko manage karega.

package websocket

type Hub struct {
	Clients map[*Client]bool

	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {

	for {
		select {

		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.Unregister:
			delete(h.Clients, client)

		case message := <-h.Broadcast:
			for client := range h.Clients {
				client.Send <- message
			}
		}
	}
}
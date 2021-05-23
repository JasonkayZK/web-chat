package server

import (
	"encoding/json"
)



var Hub = &ClientGroup{
	clients:    make(map[*Client]bool),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	broadcast:  make(chan []byte),
}

func (h *ClientGroup) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
			c.Message.MessageType = "init"
			c.Message.IP = c.Socket.RemoteAddr().String()
			c.Message.UserList = userList
			msgData, _ := json.Marshal(c.Message)
			c.Send <- msgData
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.Send)
			}
		case data := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.Send <- data:
				default:
					delete(h.clients, c)
					close(c.Send)
				}
			}
		}
	}
}

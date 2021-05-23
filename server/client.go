package server

import (
	"github.com/gorilla/websocket"
	"github.com/jasonkayzk/web-chat/model"
)

// 单个socket客户端属性
type Client struct {
	Socket  *websocket.Conn
	Send    chan []byte
	Message *model.Message
}

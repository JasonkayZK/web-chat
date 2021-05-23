package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jasonkayzk/web-chat/model"
	"net/http"
)

var (
	upGrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	userList []map[string]string
)

func WsServer(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("出错咯～", err)
	}
	// register：客户端连接，定义客户端基本结构
	client := &Client{Socket: ws, Send: make(chan []byte), Message: &model.Message{}}
	Hub.register <- client
	go client.Writer()
	client.Reader()

	// 当前 wsServer 结束时，关闭此客户端连接
	defer func() {
		client.Message.MessageType = "logout"
		userList = remove(client.Message.UUID)
		client.Message.UserList = userList
		msgData, _ := json.Marshal(client.Message)
		Hub.broadcast <- msgData /* 将客户端断开的时间写入广播的数据管道 */
		Hub.unregister <- client /* 将此客户端连接写入断开连接的数据管道 */
	}()
}

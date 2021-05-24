package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jasonkayzk/web-chat/model"
	"github.com/jasonkayzk/web-chat/service"
	"time"
)

func (c *Client) Reader() {
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Hub.unregister <- c
			break
		}
		_ = json.Unmarshal(message, &c.Message)
		handleMsgType(c)
	}
}

func (c *Client) Writer() {
	for message := range c.Send {
		if c.Message.MessageType == "private" {
			if c.Message.ToUUID == c.Message.UUID {
				_ = c.Socket.WriteMessage(websocket.TextMessage, message)
			}
		} else {
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
	_ = c.Socket.Close()
}

func handleMsgType(c *Client) {
	switch c.Message.MessageType {
	case "ping":
		return
	case "login":
		// 1：插入列表
		userList = insert(c.Message.UUID, c.Message.Username)
		c.Message.UserList = userList

		// 2：写入登录日志
		go service.AddLoginLog(&model.LoginLog{
			UUID:       c.Message.UUID,
			IP:         c.Message.IP,
			LogType:    model.Login,
			Username:   c.Message.Username,
			InsertTime: time.Now().UnixNano(),
		})

		// 3：更新聊天人数峰值
		go service.UpdateMaxUserNum(len(userList))

		// 4：发送欢迎通知
		msgData, _ := json.Marshal(c.Message)
		Hub.broadcast <- msgData
	case "message":
		// 1：消息限频
		banned, err := service.AddChatCount(c.Message.UUID)
		if err != nil {
			fmt.Println(err)
			return
		}
		if banned {
			c.Message.Content = fmt.Sprintf("%s 发言过快,请稍后重试~", c.Message.Username)
			msgData, _ := json.Marshal(c.Message)
			Hub.broadcast <- msgData
			return
		}

		// 2：add message data log
		go service.AddChatMessage(c.Message)

		// 3：广播消息
		msgData, _ := json.Marshal(c.Message)
		Hub.broadcast <- msgData
	case "private":
		// 直接转发私人消息
		msgData, _ := json.Marshal(c.Message)
		Hub.broadcast <- msgData
	case "logout":
		// 1:删除用户列表
		c.Message.MessageType = "logout"
		userList = remove(c.Message.UUID)
		c.Message.UserList = userList

		// 2:写入退出日志
		go service.AddLoginLog(&model.LoginLog{
			UUID:       c.Message.UUID,
			IP:         c.Message.IP,
			LogType:    model.Logout,
			Username:   c.Message.Username,
			InsertTime: time.Now().UnixNano(),
		})

		// 3:广播退出消息
		msgData, _ := json.Marshal(c.Message)
		Hub.broadcast <- msgData
		Hub.unregister <- c
	default:
		fmt.Println("=======================")
	}
}

// 将新加入的成员写入user_list
func insert(uuid, user string) []map[string]string {
	userList = append(userList, map[string]string{
		"uuid":     uuid,
		"username": user,
	})
	return userList
}

// 将断线的成员移除
func remove(uuid string) []map[string]string {
	var newUserList []map[string]string
	for _, item := range userList {
		if item["uuid"] != uuid {
			fmt.Println(item["uuid"], uuid)
			newUserList = append(newUserList, item)
		}
	}
	return newUserList
}

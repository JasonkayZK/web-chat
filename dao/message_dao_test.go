package dao

import (
	"fmt"
	"github.com/jasonkayzk/web-chat/model"
	"github.com/jasonkayzk/web-chat/util"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	util.InitMongo()
	m.Run()
}

func TestInsertMessage(t *testing.T) {
	for i := 0; i < 50; i++ {
		messageId, err := InsertMessage(
			&model.MessageResponse{
				UUID:        "O0I9vLCltD",
				IP:          "[::1]:51355",
				MessageType: "message",
				Username:    "嘎嘎",
				Content:     fmt.Sprintf("你好：%d", i),
				MessageTime: int(time.Now().Unix() * 1000),
			},
		)
		if err != nil {
			panic(err)
		}

		fmt.Println(messageId)
	}
}

func TestGetHistoryMessage(t *testing.T) {
	messages, err := GetHistoryMessage(9223372036854775807, 20)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(messages))
	for _, message := range messages {
		fmt.Println(message)
	}
}

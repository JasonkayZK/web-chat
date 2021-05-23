package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasonkayzk/web-chat/dao"
	"github.com/jasonkayzk/web-chat/model"
	"net/http"
	"strconv"
)

func AddChatMessage(message *model.Message) {
	_, err := dao.InsertMessage(message)
	if err != nil {
		fmt.Println(fmt.Sprintf("insert chat message err: %s", err.Error()))
	}
}

func GetHistoryMessage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	insertTime, err := strconv.ParseInt(c.Query("insert_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprintf("insert_time parse err: %s", err.Error()),
		})
		return
	}
	if insertTime <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprintf("insert_time below 0"),
		})
		return
	}

	messageCount, err := strconv.ParseInt(c.Query("message_count"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprintf("message_count parse err: %s", err.Error()),
		})
		return
	}
	if messageCount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprintf("message_count below 0"),
		})
		return
	}

	messages, err := dao.GetHistoryMessage(insertTime, messageCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": fmt.Sprintf("get history message err: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": messages,
	})
}

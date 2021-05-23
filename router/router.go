package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonkayzk/web-chat/server"
	"github.com/jasonkayzk/web-chat/service"
)

func InitRouter() *gin.Engine {
	go server.Hub.Run()
	router := gin.Default()

	router.GET("/chat_history", service.GetHistoryMessage)

	router.Any("/im", server.WsServer)
	return router
}

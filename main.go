package main

import (
	"github.com/jasonkayzk/web-chat/router"
	"github.com/jasonkayzk/web-chat/util"
)

func main() {
	util.InitMongo()
	util.InitRedis()
	//gin.SetMode(gin.ReleaseMode)
	webChat := router.InitRouter()
	_ = webChat.Run(":8008")
}

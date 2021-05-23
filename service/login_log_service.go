package service

import (
	"fmt"
	"github.com/jasonkayzk/web-chat/dao"
	"github.com/jasonkayzk/web-chat/model"
)

func AddLoginLog(log *model.LoginLog) {
	_, err := dao.InsertLoginLog(log)
	if err != nil {
		fmt.Println(fmt.Sprintf("insert login log err: %s", err.Error()))
	}
}

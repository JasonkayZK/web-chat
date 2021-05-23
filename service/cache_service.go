package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jasonkayzk/web-chat/util"
	"strconv"
	"time"
)

const (
	maxUserNumKey = `chat:max_user_num`

	accountChatCountPrefixTpl = `chat:banned:uuid:%s`

	bannedAccountTimePrefixTpl = `chat:banned:time:uuid:%s`

	initBannedTime = 15

	// 5秒钟内发言超过60次，被禁言
	checkCircle = 5

	circleMaxChat = 60
)

func UpdateMaxUserNum(userNum int) {
	maxUserNum, err := getMaxUserNum()
	if err != nil {
		fmt.Println(err)
	}
	if userNum > maxUserNum {
		_, err = util.RedisClient.Set(maxUserNumKey, userNum, 0).Result()
		if err != nil {
			fmt.Println(fmt.Sprintf("update max user num err: %s", err.Error()))
		}
	}
}

func getMaxUserNum() (int, error) {
	maxUserNum, err := util.RedisClient.Get(maxUserNumKey).Int()
	if err != nil {
		if err == redis.Nil {
			util.RedisClient.Set(maxUserNumKey, 1, 0)
		} else {
			return 0, fmt.Errorf("get max user num err: %s", err.Error())
		}
	}

	return maxUserNum, nil
}

func AddChatCount(uuid string) (bool, error) {
	chatCount, err := util.RedisClient.Get(fmt.Sprintf(accountChatCountPrefixTpl, uuid)).Result()
	if err != nil && err != redis.Nil {
		return false, fmt.Errorf("add chat count err: %s", err)
	}

	if chatCount == "" || err == redis.Nil {
		_, err = util.RedisClient.SetNX(fmt.Sprintf(accountChatCountPrefixTpl, uuid), 0, checkCircle*time.Second).Result()
		if err != nil {
			return false, fmt.Errorf("set chat count err: %s", err.Error())
		}
		return false, nil
	} else {
		_, err := util.RedisClient.Incr(fmt.Sprintf(accountChatCountPrefixTpl, uuid)).Result()
		if err != nil {
			return false, fmt.Errorf("incr chat count err: %s", err.Error())
		}

		ban, err := checkBan(uuid)
		if err != nil {
			return false, fmt.Errorf("check ban err: %s", err.Error())
		}

		if ban {
			err := banUUID(uuid)
			if err != nil {
				return false, fmt.Errorf("ban uuid err: %s", err.Error())
			}
			return true, nil
		}
		return false, nil
	}
}

func banUUID(uuid string) error {
	nextBannedTime, err := checkBannedUUID(uuid)
	if err != nil {
		return err
	}

	var bannedTime time.Duration
	if nextBannedTime > 0 {
		bannedTime = time.Duration(nextBannedTime)
	} else {
		bannedTime = initBannedTime
	}

	_, err = util.RedisClient.Set(fmt.Sprintf(accountChatCountPrefixTpl, uuid), 0, checkCircle*time.Second).Result()
	if err != nil {

	}

	_, err = util.RedisClient.Set(fmt.Sprintf(bannedAccountTimePrefixTpl, uuid), nextBannedTime*2, bannedTime*time.Second).Result()
	if err != nil {

	}

	return nil
}

func checkBan(uuid string) (bool, error) {
	chatCount, err := util.RedisClient.Get(fmt.Sprintf(accountChatCountPrefixTpl, uuid)).Result()
	if err != nil && err != redis.Nil {
		return false, fmt.Errorf("get chat count err: %s", err.Error())
	}
	if chatCount == "" || err == redis.Nil {
		return false, nil
	}

	count, err := strconv.ParseInt(chatCount, 10, 64)
	if err != nil {
		return false, fmt.Errorf("parse chat count err: %s", err.Error())
	}

	return count > circleMaxChat, nil
}

func checkBannedUUID(uuid string) (int64, error) {
	result, err := util.RedisClient.Get(fmt.Sprintf(bannedAccountTimePrefixTpl, uuid)).Result()
	if err != nil && err != redis.Nil {
		return 0, fmt.Errorf("check banned uuid err: %s", err.Error())
	}
	if result == "" || err == redis.Nil {
		return 0, nil
	}

	nextBannedTime, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse nextBannedTime err: %s", err.Error())
	}

	return nextBannedTime, nil
}

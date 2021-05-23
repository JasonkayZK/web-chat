package model

const (
	Login  = `login`
	Logout = `logout`
)

type LoginLog struct {
	UUID       string `json:"uuid" bson:"uuid"`
	IP         string `json:"ip" bson:"ip"`
	LogType    string `json:"log_type" bson:"log_type"`
	Username   string `json:"username" bson:"username"`
	InsertTime int64  `json:"insert_time" bson:"insert_time"`
}

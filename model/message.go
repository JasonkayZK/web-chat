package model

// 数据传输的数据结构
type Message struct {
	UUID        string              `json:"uuid" bson:"uuid"`
	IP          string              `json:"ip" bson:"ip"`
	ToUUID      string              `json:"to_uuid" bson:"to_uuid"`
	MessageType string              `json:"message_type" bson:"message_type"`
	Username    string              `json:"username" bson:"username"`
	Content     string              `json:"content" bson:"content"`
	MessageTime int                 `json:"message_time" bson:"message_time"`
	UserList    []map[string]string `json:"user_list" bson:"user_list"`
	InsertTime  int64               `json:"insert_time" bson:"insert_time"`
}

// 历史消息数据数据响应的数据结构
type MessageResponse struct {
	UUID        string              `json:"uuid" bson:"uuid"`
	IP          string              `json:"ip" bson:"ip"`
	ToUUID      string              `json:"to_uuid" bson:"to_uuid"`
	MessageType string              `json:"message_type" bson:"message_type"`
	Username    string              `json:"username" bson:"username"`
	Content     string              `json:"content" bson:"content"`
	MessageTime int                 `json:"message_time" bson:"message_time"`
	InsertTime  int64               `json:"insert_time" bson:"insert_time"`
}

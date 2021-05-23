## **一个简单的在线聊天室**

一个使用Golang + WebSocket实现的在线聊天室；

用到的技术栈：

-   Gin；
-   WebSocket；
-   Redis；
-   MongoDB；

在线展示：

-   [畅所欲言](https://jasonkayzk.github.io/chat/)

<br/>

### **项目特色**

-   前端使用纯静态页面（HTML + CSS + JS），即开即用；
-   后端采用MongoDB + Redis，无需创建数据库表，直接部署即可；
-   支持加载聊天历史记录；
-   使用Redis做聊天限频；

<br/>

### **使用方法**

#### **配置并启动后端**

首先，修改util目录下的`mongo.go`和`redis.go`文件：

```go
// util/mongo.go
const (
	mongoUrlTpl   = `mongodb://%s:%s@%s:%s/%s?authSource=admin&w=majority&readPreference=primary&retryWrites=true&ssl=false`
	mongoUsername = `admin`
	mongoPassword = `passwd`
	mongoHost     = `127.0.0.1`
	mongoPort     = `27017`
	mongoDBName   = `chat`
)

// util/redis.go
const (
	redisPassword = `passwd`
	redisHost     = `127.0.0.1`
	redisPort     = `6379`
)
```

修改为你的Redis和Mongo的内容；

随后启动Go项目：

```bash
go run main.go
```

成功启动后，输出相关日志：

```bash
$ go run main.go
Connected to MongoDB!
Connect to Redis!
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)
[GIN-debug] GET    /chat_history             --> github.com/jasonkayzk/web-chat/service.GetHistoryMessage (3 handlers)
[GIN-debug] GET    /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] POST   /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] PUT    /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] PATCH  /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] HEAD   /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] OPTIONS /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] DELETE /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] CONNECT /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] TRACE  /im                       --> github.com/jasonkayzk/web-chat/server.WsServer (3 handlers)
[GIN-debug] Listening and serving HTTP on :8008

```

可以看到，项目已经启动；

<br/>

#### **项目前端**

对于前端页面来说，仅仅是静态的页面，直接打开位于web目录下的`index.html`即可；

>   部署时可能需要修改`web/js/websocket.js`文件中的WebSocket地址：
>
>   ```javascript
>   uname = uname.trim()
>   ws = new WebSocket("ws://your_host:your_port/im");
>   system("正在连接服务器...")
>   ```

<br/>

#### **项目测试**

打开前端项目后，需要先填写用户的名称：

![](https://cdn.jsdelivr.net/gh/jasonkayzk/web-chat@main/images/demo_1.png)

填写后即可进入聊天室界面：

![](https://cdn.jsdelivr.net/gh/jasonkayzk/web-chat@main/images/demo_2.png)

打开多个Tab窗口并取名，可以模拟多人聊天：

![](https://cdn.jsdelivr.net/gh/jasonkayzk/web-chat@main/images/demo_3.png)

发送消息，即可进行群聊天了~

![](https://cdn.jsdelivr.net/gh/jasonkayzk/web-chat@main/images/demo_4.png)

![](https://cdn.jsdelivr.net/gh/jasonkayzk/web-chat@main/images/demo_5.png)

<br/>

当然也可以点击聊天窗口上方的`点击加载更多历史消息`，获取历史消息：

![](https://cdn.jsdelivr.net/gh/jasonkayzk/web-chat@main/images/demo_6.png)

<br/>

### **项目参考**

-   [debugxxs/gin-webim](https://github.com/debugxxs/gin-webim)


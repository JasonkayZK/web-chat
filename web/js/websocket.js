Date.prototype.format = function (fmt = "yyyy-MM-dd hh:mm:ss") {
    let o = {
        "M+": this.getMonth() + 1,                 //月份
        "d+": this.getDate(),                    //日
        "h+": this.getHours(),                   //小时
        "m+": this.getMinutes(),                 //分
        "s+": this.getSeconds(),                 //秒
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度
        "S": this.getMilliseconds()             //毫秒
    };
    if (/(y+)/.test(fmt)) {
        fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    }
    for (let k in o) {
        if (new RegExp("(" + k + ")").test(fmt)) {
            fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
        }
    }
    return fmt;
}

let uname = sessionStorage.getItem("uname");
let uuid = sessionStorage.getItem("uuid");
let insertTime = Date.now() * 1000000;
let messageCount = 20;

if (isEmpty(uname) || isEmpty(uuid)) {
    uname = prompt('给自己起个响亮的名字吧');
    uuid = createUUID(10);

    sessionStorage.setItem("uname", uname);
    sessionStorage.setItem("uuid", uuid);
}

let ws = undefined;
let wsUrl = "ws://localhost:8008/im";
let sendFixHeartTimer = null;
const heartBeatTime = 1000;

if (uname) {
    uname = uname.trim()
    ws = new WebSocket(wsUrl);
    system("正在连接服务器...")

    ws.onopen = function () {
        ws.send(JSON.stringify({
            "message_type": "login",
            "uuid": uuid,
            "content": "Hello Go WebSocket",
            "username": uname
        }));
        sendFixHeartBeat();
    };

    ws.onmessage = function (evt) {// 绑定收到消息事件
        console.log("Received Message: " + evt.data);
        const data = JSON.parse(evt.data)
        showMessage(data);
    };

    ws.onclose = function () { //绑定关闭或断开连接事件
        system("与服务器连接断开", "error")
        ws.send(JSON.stringify({
            "message_type": "logout",
            "uuid": uuid,
            "content": "下线",
            "username": uname,
            "message_time": new Date().getTime()
        }))
        clearInterval(sendFixHeartTimer)
    };
} else {
    system("服务器未连接，请给自己起个名字吧～，<a href=''>点我起名</a>")
}

document.onkeydown = function (event) {
    let e = event || window.event;
    if (e && e.keyCode === 13) {
        e.preventDefault()
        send();
    }
};

// 显示消息
function showMessage(data) {
    switch (data.message_type) {
        case "init":
            system(`服务器连接成功`, "error")
            system("欢迎来到在线 IM 聊天室；在这里你可以畅所欲言，但严禁发送违法、诈骗等信息！", "error")
            break;
        case "login":
            userListDom(data.user_list)
            system(`${data.username} 进入群聊`, "success")
            break;
        case "message":
            var {username, content} = data
            const message = `<span style="color: #40a9ff;">${new Date().format()} ${username}: </span>${content}`
            acceptMessage(message)
            break;
        case "private":
            var {username, content, to_uuid} = data
            if (to_uuid === uuid) {
                const message = `<span style="color: red;">${new Date().format()} ${username}: </span>对你说 ${content}`
                acceptMessage(message)
            }
            break;
        case "logout":
            userListDom(data.user_list)
            system(`${data.username} 已下线`, "error")
            break;
    }
}

// 发送消息
function send() {
    const message = document.getElementById("content-value").value
    if (message.trim().toString().length <= 0) {
        alert("请输入发送的内容")
        return
    }
    const UUID = document.getElementById("UUID").value;
    const message_type = UUID ? "private" : "message"
    if (message_type === "private") {
        const info = `${uname}: ${message}`
        acceptMessage(info)
    }
    ws.send(JSON.stringify({
        "message_type": message_type,
        "content": message.trim(),
        "username": uname,
        "to_uuid": UUID,
        "message_time": new Date().getTime()
    }));
    document.getElementById("UUID").value = "";
    document.getElementById("content-value").value = "";
    backToButton();
}

// 接收消息
function acceptMessage(message) {
    document.getElementById("content").innerHTML += `
                <div style="line-height: 30px;">${message}</div>
            `;
    // setTimeout(() => {
    //     document.getElementById("content").scrollTo(0, document.getElementById("content").offsetHeight);
    // }, 1000)
}

// 接收历史消息
function acceptHistoryMessage(message) {
    let para = document.createElement("div");
    para.textContent = `${message}`;
    para.style.lineHeight = '30px'

    document.getElementById("content-header").insertAdjacentHTML("afterend", `
                <div style="line-height: 30px;">${message}</div>
           `,);
}

// 系统消息通知
function system(message, message_type = "loading") {
    document.getElementById("content").innerHTML += `
        <div class="system ${message_type}"><span>${new Date().format()} </span>系统消息：${message}</div>`
}

function userListDom(userList) {
    document.getElementById("user-list").innerHTML = ""
    userList.map(item => {
        console.log(uuid, item.uuid)
        if (uuid === item.uuid) {
            document.getElementById("user-list").innerHTML += `<li style="color: red;">${item.username}(我)</li>`
        } else {
            document.getElementById("user-list").innerHTML += `
                        <li onclick="privateMessage('${item.username}', '${item.uuid}')">${item.username}</li>
                    `
        }
    })
}

// 获取历史群聊消息
function showChatHistory() {
    let httpRequest = new XMLHttpRequest();
    let url = `http://localhost:8008/chat_history?insert_time=${insertTime}&message_count=${messageCount}`;

    httpRequest.open('GET', url, true);
    httpRequest.send();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.status === 200 && httpRequest.readyState === 4) {
            let jsonText = httpRequest.responseText;
            let data = JSON.parse(jsonText).data;
            if (data === null) {
                const message = `<div class="system error"><span>${new Date().format()} </span>系统消息：已无历史聊天记录~</div>`
                acceptHistoryMessage(message);
                return;
            }
            for (let i = 0; i < data.length; i++) {
                insertTime = Math.min(data[i].insert_time, insertTime);
                const message = `<span style="color: #40a9ff;">${new Date(data[i].message_time).format()} </span><span>${data[i].username}: </span>${data[i].content}`
                acceptHistoryMessage(message);
            }
            console.log(insertTime)
        }
    };
}

function privateMessage(user, uuid) {
    document.getElementById("UUID").value = uuid
    document.getElementById("content-value").value = `@${user} `
}

function createUUID(len, radix = null) {
    let chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
    let uuid = [], i;
    radix = radix || chars.length;
    if (len) {
        for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random() * radix];
    } else {
        let r;
        uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
        uuid[14] = '4';
        for (i = 0; i < 36; i++) {
            if (!uuid[i]) {
                r = 0 | Math.random() * 16;
                uuid[i] = chars[(i === 19) ? (r & 0x3) | 0x8 : r];
            }
        }
    }
    return uuid.join('');
}

function isEmpty(obj) {
    return typeof obj === 'undefined' || obj == null || obj === '';
}

function backToButton() {
    let content = document.getElementById("content");
    content.scrollTop = content.scrollHeight;
}

function sendFixHeartBeat() {
    if (ws !== null) {
        clearInterval(sendFixHeartTimer);
        sendFixHeartTimer = setInterval(() => {
            ws.send(JSON.stringify({
                "message_type": "ping",

            }))
        }, heartBeatTime);
    }
}

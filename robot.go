package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	yaml "gopkg.in/yaml.v2"
)

const ConfigPath = "config.yaml" //配置文件

var config Config
var api openapi.OpenAPI
var ctx context.Context

// 配置加载
func init() {
	content, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		log.Println("读取配置文件出错，err = ", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Println("解析配置文件出错，err = ", err)
		os.Exit(1)
	}
	if flag == "all" || flag == "config" {
		DPrintf("配置文件为:%v", config)
	}
}

func main() {
	token := token.BotToken(config.AppID, config.Token)
	api = botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
	ctx = context.Background()
	ws, err := api.WS(ctx, nil, "") //websocket
	if err != nil {
		log.Fatalln("websocket错误， err = ", err)
		os.Exit(1)
	}

	var atMessage event.ATMessageEventHandler = atMsghandler

	intent := websocket.RegisterHandlers(atMessage)     // 注册socket消息处理
	botgo.NewSessionManager().Start(ws, token, &intent) // 启动socket监听
}

//处理 @机器人的事件
func atMsghandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	res := message.ParseCommand(data.Content)
	if flag == "all" || flag == "cmd" {
		DPrintf("cmd解析:%v", res)
	}
	cmd := res.Cmd
	// content := res.Content
	switch cmd {
	case "/hello":
		api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "你好"})
	}
	return nil
}

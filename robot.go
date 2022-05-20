package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
var citynames map[string]string
var db Storerage

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

	citynames = make(map[string]string)
	getAllCityData(citynames)

	db = &MemKV{KV: make(map[string][]string)}
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
	content := res.Content
	switch cmd {
	case "/帮助":
		switch content {
		case "":
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: ShowDefaultManul("")})
		case "世界时间":
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: ShowCommandManul("世界时间")})
		case "当前时间":
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: ShowCommandManul("当前时间")})
		case "添加日志":
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: ShowCommandManul("添加日志")})
		case "日志查询":
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: ShowCommandManul("日志查询")})
		case "日志删除":
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: ShowCommandManul("日志删除")})
		default:
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: ShowDefaultManul("请输入正确指令！！！")})
		}
	case "/世界时间":
		globaltime := getTimeofGlobalCity(content)
		if globaltime != nil {
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: CreateArkByGlobalTime(globaltime)})
		} else {
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: CreateSuccessArk("请正确输入城市名称！！！")})
		}
	case "/当前时间":
		curtime := getNowTimeofBEIJIN()
		if curtime != nil {
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: CreateArkByCurrentTime(curtime)})
		}
	case "/添加日志":
		key, val := CreateKVforStore(data.Author.ID, content)
		err := db.Put(key, val)
		if err != nil {
			log.Println("添加日志错误，err = ", err)
			return nil
		}
		api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: CreateSuccessArk("成功添加！！！")})
	case "/日志查询":
		s := strings.Split(content, " ")
		key, err := CreateKeyforQuery(data.Author.ID, s[0])
		DPrintf("%v", key)
		if err != nil {
			log.Println("日志查询错误，err = ", err)
			return nil
		}
		val, _ := db.Get(key)
		if flag == "all" || flag == "log" {
			DPrintf("%v", val)
		}
		directMsg, _ := api.CreateDirectMessage(ctx, &dto.DirectMessageToCreate{
			SourceGuildID: data.GuildID,
			RecipientID:   data.Author.ID,
		})
		if len(s) > 1 && s[1] == "公开" {
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: CreateQueryResult(val, s[0])})
		} else {
			api.PostDirectMessage(ctx, directMsg, &dto.MessageToCreate{MsgID: data.ID, Ark: CreateQueryResult(val, s[0])})
		}
	case "/日志删除":
		key, err := CreateKeyforQuery(data.Author.ID, content)
		if err != nil {
			log.Println("创建Key出错，err = ", err)
			return nil
		}
		err = db.Del(key)
		if err != nil {
			log.Println("db删除Key = ", key, " 出错，err = ", err)
			return nil
		}
		api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: CreateSuccessArk("成功删除！！！")})
	case "/撤回":
	case "/计时":
	case "/提醒":
	default:
		api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: CreateSuccessArk("请输入正确指令！")})
	}
	return nil
}

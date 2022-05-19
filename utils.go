package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/tencent-connect/botgo/dto"
)

// Debugging
const Debug = true
const flag = "all"

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}

//NewArk 创建Ark，用来简化ark创建
func NewArk(id int, kv ...*dto.ArkKV) *dto.Ark {
	return &dto.Ark{
		TemplateID: id,
		KV:         kv,
	}
}

//NewArkKV 创建ArkKV，表示key-value
func NewArkKV(key string, value string) *dto.ArkKV {
	return &dto.ArkKV{
		Key:   key,
		Value: value,
	}
}

//NewArkKObj 创建ArkKV，表示key-obj
func NewArkKObj(key string, objKv ...*dto.ArkObjKV) *dto.ArkKV {
	return &dto.ArkKV{
		Key: key,
		Obj: NewArkObj(objKv),
	}
}

//NewArkObj 创建ArkObj的数组
func NewArkObj(objKv []*dto.ArkObjKV) []*dto.ArkObj {
	array := make([]*dto.ArkObj, len(objKv))
	for i := 0; i < len(array); i++ {
		array[i] = &dto.ArkObj{
			ObjKV: []*dto.ArkObjKV{
				objKv[i],
			},
		}
	}
	return array
}

//NewArkObjKV 创建ArkObjKV
func NewArkObjKV(key string, value string) *dto.ArkObjKV {
	return &dto.ArkObjKV{
		Key:   key,
		Value: value,
	}
}

//创建当前时间的ark消息
func CreateArkByCurrentTime(curtime *TimeResp) *dto.Ark {
	return NewArk(23,
		NewArkKV("#DESC#", "描述"),
		NewArkKV("#PROMPT#", "提示信息"),
		NewArkKObj("#LIST#",
			NewArkObjKV("desc", curtime.Data.Datetime),
			NewArkObjKV("desc", curtime.Data.Week)))
}

//创建世界时间的ark消息
func CreateArkByGlobalTime(globaltime *GlobaltimeResp) *dto.Ark {
	return NewArk(23,
		NewArkKV("#DESC#", "描述"),
		NewArkKV("#PROMPT#", "提示信息"),
		NewArkKObj("#LIST#",
			NewArkObjKV("desc", globaltime.Data.Continent+" "+globaltime.Data.Country+" "+globaltime.Data.City+" "+globaltime.Data.Time_zone),
			NewArkObjKV("desc", "当地时间："+globaltime.Data.Datetime+" "+globaltime.Data.Week),
			NewArkObjKV("desc", "北京时间："+globaltime.Data.BJ_datetime)))
}

//判断字符串是否全是英文
func StrAllLetter(str string) bool {
	match, _ := regexp.MatchString(`^[A-Za-z]+$`, str)
	return match
}

//创建当天需要存储的kv对
//使用userid + 当前日期作为key，中间加" "作为分隔，便于后期切分
//使用当前时间+content作为val，中间加:,便于后面输出日志排版
func CreateKVforStore(user string, content string) (key string, val string) {
	curtime := getNowTimeofBEIJIN()
	if flag == "all" || flag == "log" {
		DPrintf("%v+%d", curtime.Data.Datetime, len(curtime.Data.Datetime))
	}
	s := strings.Split(curtime.Data.Datetime, " ")
	if flag == "all" || flag == "log" {
		DPrintf("%v+%v+%d", s[0], s[1], len(s))
	}
	key = user + " " + s[0]
	val = s[1] + ": " + content
	if flag == "all" || flag == "log" {
		DPrintf("{key:%v},{val:%v}", key, val)
	}
	return key, val
}

//创建成功的Ark
func CreateSuccessArk(msg string) *dto.Ark {
	return NewArk(23,
		NewArkKV("#DESC#", "描述"),
		NewArkKV("#PROMPT#", "提示信息"),
		NewArkKObj("#LIST#",
			NewArkObjKV("desc", msg)))
}

//创建用于查询的key
func CreateKeyforQuery(user string, content string) (string, error) {
	key := user + " " + content
	return key, nil
}

//创建查询结果的Ark
func CreateQueryResult(val []string, date string) *dto.Ark {
	logs := make([]*dto.ArkObjKV, len(val)+1)
	logs[0] = NewArkObjKV("desc", date+"日志")
	for i, v := range val {
		logs[i+1] = NewArkObjKV("desc", v)
	}
	return NewArk(23, NewArkKObj("#LIST#", logs...))
}

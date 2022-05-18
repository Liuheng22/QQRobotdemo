package main

import (
	"log"

	"github.com/tencent-connect/botgo/dto"
)

// Debugging
const Debug = false
const flag = ""

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

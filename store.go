package main

import "fmt"

type Storerage interface {
	//获取对应key的所有数据
	Get(key string) ([]string, error)
	//添加对应的kv对
	Put(key string, value string) error
	//删除对应key的所有数据
	Del(key string) error
}

//内存kv存储器，存用户上传的日志
type MemKV struct {
	KV map[string][]string
}

func (kv *MemKV) Get(key string) ([]string, error) {
	val, ok := kv.KV[key]
	if !ok {
		return nil, fmt.Errorf("不存在相应数据")
	}
	return val, nil
}
func (kv *MemKV) Put(key string, value string) error {
	_, ok := kv.KV[key]
	if ok {
		kv.KV[key] = append(kv.KV[key], value)
		return nil
	}
	values := make([]string, 0)
	values = append(values, value)
	kv.KV[key] = values
	return nil
}
func (kv *MemKV) Del(key string) error {
	_, ok := kv.KV[key]
	if ok {
		kv.KV[key] = nil
	}
	return nil
}

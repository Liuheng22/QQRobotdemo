package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//获取当前的北京时间
func getNowTimeofBEIJIN() *TimeResp {
	url := "http://api.k780.com/?app=life.time&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json"
	body := getNewworkData(url)
	if body == nil {
		return nil
	}
	var curtime TimeResp
	err := json.Unmarshal(body, &curtime)
	if err != nil {
		log.Fatalln("数据解析异常，err = ", err, body)
		return nil
	}
	if curtime.Success != "1" {
		log.Fatalln("返回数据失败，err = ", curtime.Msg)
		return nil
	}
	return &curtime
}

//获得网络请求数据
func getNewworkData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("数据接口请求异常，err = ", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("数据接口请求异常，err = ", err)
		return nil
	}
	return body
}

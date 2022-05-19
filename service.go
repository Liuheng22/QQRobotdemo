package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//将中文名字转换为英文名字
func getCityCn2CityEn(city string) string {
	url := "https://way.jd.com/he/freecity?city=" + city + "&appkey=da39dce4f8aa52155677ed8cd23a6470"
	body := getNewworkData(url)
	if body == nil {
		return ""
	}
	DPrintf("body:%v", string(body))
	var citydata CityDataresp
	err := json.Unmarshal(body, &citydata)
	if err != nil {
		log.Fatalln("数据解析异常，err = ", err, body)
		return ""
	}
	if flag == "all" || flag == "parser" {
		DPrintf("城市数据Json解析:%v", citydata)
	}
	if citydata.Success != "10000" {
		log.Fatalln("返回数据失败，err = ", citydata.Msg)
		return ""
	}
	if flag == "all" || flag == "parser" {
		DPrintf("城市数据Json解析:%v", citydata)
	}
	city_en := citydata.Data.Hew.City_En
	return city_en
}

//获得所有城市信息
func getAllCityData(citynames map[string]string) {
	url := "http://api.k780.com/?app=time.world_city&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json"
	body := getNewworkData(url)
	if body == nil {
		return
	}
	var citydatas CitysData
	err := json.Unmarshal(body, &citydatas)
	if err != nil {
		log.Fatalln("数据解析异常，err = ", err, body)
		return
	}
	if citydatas.Success != "1" {
		log.Fatalln("返回数据失败")
		return
	}
	for _, cityname := range citydatas.Result.Data {
		citynames[cityname.CityCn] = cityname.CityEn
	}
	return
}

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

//获得全球时间
func getTimeofGlobalCity(city string) *GlobaltimeResp {
	city_en := city
	if !StrAllLetter(city) {
		// city_en = getCityCn2CityEn(city)
		city_en = citynames[city]
	}
	url := "http://api.k780.com/?app=time.world&city_en=" + city_en + "&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json"
	body := getNewworkData(url)
	if body == nil {
		return nil
	}
	var globaltime GlobaltimeResp
	err := json.Unmarshal(body, &globaltime)
	if err != nil {
		log.Println("数据解析异常，err = ", err, body)
		return nil
	}
	if globaltime.Success != "1" {
		log.Println("返回数据失败，err = ", globaltime.Msg)
		return nil
	}
	return &globaltime
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

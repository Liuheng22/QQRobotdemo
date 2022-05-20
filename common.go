package main

// Config 定义了配置文件的结构
type Config struct {
	AppID uint64 `yaml:"appid"` //机器人的appid
	Token string `yaml:"token"` //机器人的token
}

//当前时间的resp
type TimeResp struct {
	Success string `json:"success"` //标识请求是否成功，1表示成功，0表示失败
	Data    Time   `json:"result"`  //请求成功时获得的数据
	Msg     string `json:"msg"`     //请求失败时获得的原因
}

type Time struct {
	Datetime string `json:"datetime_2"` //日期，如2016年06月24日 15时40分49秒
	Week     string `json:"week_2"`     //星期几，如星期五
}

//世界时间的resp
type GlobaltimeResp struct {
	Success string     `json:"success"` //标识请求是否成功，1表示成功，0表示失败
	Data    Globaltime `json:"result"`  //请求成功时获得的数据
	Msg     string     `json:"msg"`     //请求失败时获得的原因
}

type Globaltime struct {
	Continent   string `json:"continents_cn"` //所属大洲
	Country     string `json:"contry_cn"`     //所属国家
	City        string `json:"city_cn"`       //城市名
	Time_zone   string `json:"time_zone_nm"`  //时区名
	Datetime    string `json:"datetime_1"`    //当地时间
	Week        string `json:"week_2"`        //周几
	BJ_datetime string `json:"bjt_datetime"`  //北京时间
}

//城市查询的resp
type CityDataresp struct {
	Success string   `json:"code"`   //状态码，10000为成功
	Data    CityData `json:"result"` //请求成功时获得的数据
	Msg     string   `json:"msg"`    //请求回复的消息
}

type CityData struct {
	Hew HeWeather `json:"HeWeater5"` //城市基本信息
}

type HeWeather struct {
	City_En string `json:"city"` //城市英文名
}

//世界城市查询resp
type CitysData struct {
	Success string     `json:"success"` //请求是否成功
	Result  CityResult `json:"result"`  //世界城市查询结果
}

type CityResult struct {
	Data map[string]CityName `json:"lists"` //所有城市信息
}

type CityName struct {
	CityCn string `json:"cityCn"` //城市中文名
	CityEn string `json:"cityEn"` //城市英文名
}

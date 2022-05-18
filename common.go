package main

// Config 定义了配置文件的结构
type Config struct {
	AppID uint64 `yaml:"appid"` //机器人的appid
	Token string `yaml:"token"` //机器人的token
}

type TimeResp struct {
	Success string `json:"success"` //标识请求是否成功，1表示成功，0表示失败
	Data    Time   `json:"result"`  //请求成功时获得的数据
	Msg     string `json:"msg"`     //请求失败时获得的原因
}

type Time struct {
	Datetime string `json:"datetime_2"` //日期，如2016年06月24日 15时40分49秒
	Week     string `json:"week_2"`     //星期几，如星期五
}

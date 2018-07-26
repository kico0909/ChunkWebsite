package config

import (
	"io/ioutil"
	"log"
	"utils/ChunkLib/fileSystem"
	"encoding/json"
)

type tlsData struct {
	Open bool `json:"open"`
	Domain string `json:"domain"`
	Email string `json:"email"`
}
type mysqlSetOpt struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
	Dbname string `json:"dbname"`
	Socket string `json:"socket"`
}
type mysqlData struct {
	Key bool `json:"key"`
	Default mysqlSetOpt `json:"default"`
}
type RedisData struct {
	Key bool `json:"key"`
	Host string `json:"host"`
	Port string `json:"port"`
	Dbname string `json:"dbname"`
	Password string `json:"password"`
}
type CasData struct {
	Key bool `json:"key"`
	Server string `json:"server"`
	WhiteList []string `json:"whiteList"`
}

type SessionOpt struct {
	Key bool `json:"key"`
	SessionType string `json:"sessionType"`
	SessionName string `json:"sessionName"`
	SessionLifeTime int64 `json:"sessionLifeTime"`
	SessionRedis RedisData `json:"sessionRedis"`
}

type ConfigData struct {
	WebTitle string `json:"webTitle"`
	TemplateUrl string `json:"templateUrl"`
	StaticFilePath string `json:"staticFilePath"`
	IsAllStatic bool `json:"isAllStatic"`
	WebPort int64 `json:"webPort"`
	TLS tlsData `json:"tls"`
	Mysql mysqlData `json:"mysql"`
	Redis RedisData `json:"redis"`
	Session SessionOpt `json:"session"`
	Cas CasData `json:"cas"`
	Custom map[string]interface{} `json:"custom"`
}

func Init() ConfigData{
	log.Println("读取配置文件 ...\n")

	var ret ConfigData

	filePath, err := fileSystem.GetMyPath()

	if err!= nil {
		log.Fatal(err)
	}

	cont, err := ioutil.ReadFile(filePath + "/conf.json")

	if err!=nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(cont, &ret); err != nil {
		log.Fatal(err)
	}

	return ret
}
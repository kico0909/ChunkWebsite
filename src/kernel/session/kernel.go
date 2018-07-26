package session

import (
	"kernel/public"
	"github.com/astaxie/beego/session"
	"encoding/json"
	"log"
	"strconv"
)

var sessionEndName = "_glsessn_"

func Init(){

	if public.WebSiteConfig.IsAllStatic && !public.WebSiteConfig.Session.Key {
		return
	}

	// 配置信息检测容错设置默认值
	if public.WebSiteConfig.Session.SessionType == "" {
		public.WebSiteConfig.Session.SessionType = "memory"
	}

	if public.WebSiteConfig.Session.SessionName == "" {
		public.WebSiteConfig.Session.SessionName = "_CHUNK"
	}

	if public.WebSiteConfig.Session.SessionLifeTime == 0 {
		public.WebSiteConfig.Session.SessionLifeTime = 3600
	}

	// 初始化 session
	sessionSetup := ""
	switch public.WebSiteConfig.Session.SessionType{

	case "redis":
		log.Print("初始化SESSION [ redis ] ! \n\n")
		srHost := public.WebSiteConfig.Session.SessionRedis.Host
		srPort := public.WebSiteConfig.Session.SessionRedis.Port
		srNumber := public.WebSiteConfig.Session.SessionRedis.Dbname
		srPassword := public.WebSiteConfig.Session.SessionRedis.Password
		sessionSetup = `{"cookieName":"` + public.WebSiteConfig.Session.SessionName + sessionEndName+`","gclifetime":`+(strconv.FormatInt(public.WebSiteConfig.Session.SessionLifeTime, 10))+`,"enableSetCookie":true,"ProviderConfig":"`+srHost+`:`+srPort+`,`+srNumber+`,`+srPassword+`"}`
		break;

	default:
		log.Print("初始化SESSION [ memory ] \n\n")
		sessionSetup = `{"cookieName":"`+public.WebSiteConfig.Session.SessionName+sessionEndName+`","gclifetime":`+(strconv.FormatInt(public.WebSiteConfig.Session.SessionLifeTime, 10))+`,"enableSetCookie":true}`
	}


	if len(sessionSetup) > 0 {
		var jsonRes session.ManagerConfig
		json.Unmarshal([]byte(sessionSetup), &jsonRes)
		public.Session, _ = session.NewManager(public.WebSiteConfig.Session.SessionType, &jsonRes)
		go public.Session.GC()
	}
}

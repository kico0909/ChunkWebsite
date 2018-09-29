package session

import (
	"kernel/public"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego/session"
	"log"
)

var sessionEndName = "_glsessn_"

var sessionSetup session.ManagerConfig

func Init(){

	var err error

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

	sessionSetup.CookieName = public.WebSiteConfig.Session.SessionName + sessionEndName
	sessionSetup.Gclifetime = public.WebSiteConfig.Session.SessionLifeTime
	sessionSetup.EnableSetCookie = true

	// 初始化 session
	switch public.WebSiteConfig.Session.SessionType{

	case "redis":
		log.Println("初始化SESSION [ redis ] !")
		srHost := public.WebSiteConfig.Session.SessionRedis.Host
		srPort := public.WebSiteConfig.Session.SessionRedis.Port
		srNumber := public.WebSiteConfig.Session.SessionRedis.Dbname
		srPassword := public.WebSiteConfig.Session.SessionRedis.Password
		sessionSetup.ProviderConfig = srHost+`:`+srPort+`,`+srNumber+`,`+srPassword
		break

	default:
		log.Println("初始化SESSION [ memory ]")
	}


	public.Session, err = session.NewManager(public.WebSiteConfig.Session.SessionType, &sessionSetup)

	if err != nil {
		log.Fatalln(err)
	}
	go public.Session.GC()

}
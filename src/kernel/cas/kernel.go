package cas

/*
服务器的内核 用于连接CAS验证服务器
*/

import (
	"kernel/public"
	"middleWare"
	"kernel/route"
	"log"
)

func Init(){

	if public.WebSiteConfig.Cas.Key && !public.WebSiteConfig.IsAllStatic && public.WebSiteConfig.Session.Key {

		// 创建 CAS 验证
		log.Print("初始化集中式权限验证服务CAS [ "+public.WebSiteConfig.Cas.Server+" ] ... \n\n")
		middleWare.Cas.Set( public.WebSiteConfig.Cas.Server, public.WebSiteConfig.Session.SessionLifeTime, "_base_user_infos_")
		route.Interceptor.AddFunc(middleWare.Cas.IsLogined)

		if len(public.WebSiteConfig.Cas.WhiteList)>0{
			log.Print("创建拦截器白名单 ... \n\n")
			route.Interceptor.AddWhiteList(public.WebSiteConfig.Cas.WhiteList)
		}

	}


}
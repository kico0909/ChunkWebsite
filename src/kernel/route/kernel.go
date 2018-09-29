package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"kernel/public"
	"routeConfig"
	"ChunkLib/fileSystem"
	"ChunkLib/cas"
	)

// 路由
var RouteRule = mux.NewRouter()

type RouterType interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// 路由初始化
func Init () RouterType {
	routePath := "/"

	// 动态路由部分
	if !public.WebSiteConfig.IsAllStatic {
		routePath = "/static/"
		log.Println("初始化动态路由 ...")
		routeConfig.Init(RouteRule)
	}

	// 静态路由
	log.Println("初始化静态路由 ...")
	systemPath, _ := fileSystem.GetMyPath()
	perfectPath := systemPath + public.WebSiteConfig.StaticFilePath
	RouteRule.PathPrefix(routePath).Handler(http.StripPrefix(routePath, http.FileServer(http.Dir(perfectPath))))

	// 创建404
	RouteRule.NotFoundHandler = http.HandlerFunc(notFound)

	// 是否添加CAS
	if public.WebSiteConfig.Cas.Key && !public.WebSiteConfig.IsAllStatic && public.WebSiteConfig.Session.Key {

		// 初始化 CAS 验证
		log.Println("初始化集中式权限验证服务CAS [ "+public.WebSiteConfig.Cas.Server+" ] ... ")
		casRouter := cas.NewCas( public.WebSiteConfig.Cas.Server, public.WebSiteConfig.Session.SessionLifeTime, "_base_user_infos_", public.Session)

		if len(public.WebSiteConfig.Cas.WhiteList)>0{
			log.Println("创建拦截器白名单 ...")
			casRouter.AddWhiteList(public.WebSiteConfig.Cas.WhiteList)
		}

		public.Cas = casRouter

		return RouterType(casRouter.Router(RouteRule))

	}else{

		return RouterType(RouteRule)
	}
}



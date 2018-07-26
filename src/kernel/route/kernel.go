package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"kernel/public"
	"routeConfig"
	"utils/ChunkLib/fileSystem"
)

// 路由
var RouteRule = mux.NewRouter()

// 拦截器
var Interceptor InterceptorClass

// 路由初始化
func Init () {
	routePath := "/"

	// 动态路由部分
	if !public.WebSiteConfig.IsAllStatic {
		routePath = "/static/"
		log.Print("初始化动态路由 ... \n\n")
		routeConfig.Init(RouteRule)
	}

	// 静态路由
	log.Print("初始化静态路由 ... \n\n")
	systemPath, _ := fileSystem.GetMyPath()
	perfectPath := systemPath + public.WebSiteConfig.StaticFilePath
	RouteRule.PathPrefix(routePath).Handler(http.StripPrefix(routePath, http.FileServer(http.Dir(perfectPath))))

	// 创建404
	RouteRule.NotFoundHandler = http.HandlerFunc(notFound)

	// 拦截器 启动
	Interceptor.New(RouteRule)
}



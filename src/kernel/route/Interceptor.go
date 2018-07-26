package route

// 路由的拦截器

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"regexp"
	"strings"
)

// 拦截器
type InterceptorClass struct {
	muxServeHTTP func(w http.ResponseWriter, r *http.Request)
	runFuncs []func(w http.ResponseWriter, r *http.Request)bool
	whiteList []*regexp.Regexp
}

func (_self *InterceptorClass) New(router *mux.Router)*InterceptorClass{
	_self.muxServeHTTP = router.ServeHTTP
	return _self
}

func (_self *InterceptorClass) AddFunc(f func(w http.ResponseWriter, r *http.Request)bool){
	_self.runFuncs = append(_self.runFuncs, f)
}

func (_self *InterceptorClass) ServeHTTP (w http.ResponseWriter, r *http.Request){
	key := true

	// 符合白名单 直接执行
	for _, reg:= range _self.whiteList {
		if reg.Match([]byte(r.URL.Path)) {
			_self.muxServeHTTP(w,r)
			return
		}
	}

	// 白名单中没有 循环执行拦截器
	for _,v := range _self.runFuncs {
		if key = v(w,r); !key {
			break
		}
	}

	if key {
		_self.muxServeHTTP(w,r)
	}
}

func (_self *InterceptorClass) AddWhiteList(urls []string){
	for _,v := range urls {
		v = strings.Replace( v,"*", "[a-z|A-Z|0-9]*", 100)
		reg, err := regexp.Compile(v)
		if err != nil {
			log.Print("拦截器白名单==>",err)
		}
		_self.whiteList = append(_self.whiteList, reg)
	}
}

func (_self *InterceptorClass) ClearWhiteList(){

}





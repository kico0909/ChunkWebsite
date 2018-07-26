package middleWare

import (
	"net/http"
)

// 检测是否是POST访问的中间件
func IsPost(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Redirect( w, r,"/404", http.StatusFound)
	}
}

// 检测是否是POST访问的中间件
func IsGet(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Redirect( w, r,"/404", http.StatusFound)
	}
}

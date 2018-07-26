package routeConfig

import (
	"app/view"
	"github.com/gorilla/mux"
)

func Init (Route *mux.Router){

	// 首页
	Route.HandleFunc("/", view.Index)
	Route.HandleFunc("/index", view.Index)

	Route.HandleFunc("/hehe/{id}/{cd}", view.Index)

	// 登录页
	Route.HandleFunc("/login", view.Login)

	// 登出页
	Route.HandleFunc("/logout", view.Logout)

}
package router

import (
	"fmt"
	"github.com/Cgo"
	"github.com/Cgo/route"
	"github.com/Cgo/funcs"
	)


func init (){

	//Cgo.Router.Filter("beforeRoute", tmpFilter)
	//Cgo.Router.Filter("afterRoute", tmpFilter)
	//Cgo.Router.Filter("afterRender", tmpFilter)

	Cgo.Router.Register("/", index).Mehtods("get")
	Cgo.Router.Register("/t1/{id}", tmp1).Mehtods("post","get")
	Cgo.Router.Register("/t2/{id}", tmp2).Mehtods("post")
	Cgo.Router.Register("/t3/ddd", tmp3).Mehtods("get")
}

func tmp1(h *route.RouterHandler){
	fmt.Fprintf(h.W,"router1")
}

func tmp2(h *route.RouterHandler){
	fmt.Fprintf(h.W,"router2")
}

func tmp3(h *route.RouterHandler){
	fmt.Fprintf(h.W,"router3")
}

func tmpFilter(h *route.RouterHandler){
	fmt.Fprint(h.W, "<br>...........filter")
}

func index(h *route.RouterHandler){
	funcs.RenderHtml(h.W, "index", nil)
}
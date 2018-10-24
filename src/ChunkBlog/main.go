package main

import (
	"github.com/Cgo"
	_ "ChunkBlog/router"	// 设置路由
	)

func main(){

	Cgo.Config.Set("./conf.json")
	Cgo.Run()
}
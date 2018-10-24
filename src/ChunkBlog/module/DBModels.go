package module

import (
	"Cgo"
)

// boss 数据模型
var Boss Cgo.Module = Cgo.Module{"boss",  Cgo.Mysql}

// 用户表模型
var User Cgo.Module = Cgo.Module{"users", Cgo.Mysql}



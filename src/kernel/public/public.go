package public

import (
	// session 的引用, 只支持 memory,cookie,redis,file
	"github.com/astaxie/beego/session"
	_"github.com/astaxie/beego/session/redis"
	"ChunkLib/redis"
	"ChunkLib/mysql"
	"kernel/config"
	"ChunkLib/cas"
)

// 网站基本配置
var WebSiteConfig config.ConfigData

// session 全局变量
var Session *session.Manager

// Redis
var Redis reids.DatabaseRedis

// Mysql
var Mysql mysql.DatabaseMysql

// Cas
var Cas cas.CasRouter


package mysql

import (
	"kernel/public"
	"log"
)

func Init () {


	if !public.WebSiteConfig.IsAllStatic && public.WebSiteConfig.Mysql.Key {

		// 默认 sql 设置
		log.Print("初始化MYSQL [ default ] 的链接 ... \n")
		public.Mysql.Init(
			"default",
			public.WebSiteConfig.Mysql.Default.Username,
			public.WebSiteConfig.Mysql.Default.Password,
			public.WebSiteConfig.Mysql.Default.Host,
			public.WebSiteConfig.Mysql.Default.Port,
			public.WebSiteConfig.Mysql.Default.Dbname,
			public.WebSiteConfig.Mysql.Default.Socket)

	}
}
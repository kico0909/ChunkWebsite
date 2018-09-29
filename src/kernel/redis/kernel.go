package redis

import (
"kernel/public"
"log"
"encoding/json"
)

func Init () {
	// 纯静态网站的逻辑判断和执行
	if  !public.WebSiteConfig.IsAllStatic && public.WebSiteConfig.Redis.Key {

		// DB redis 初始化
		log.Println("初始化 [ redis ] 库 ...")

		str, _ := json.Marshal(public.WebSiteConfig.Redis)
		tmp_map := make(map[string]interface{})
		json.Unmarshal(str, &tmp_map)
		public.Redis.Init(public.WebSiteConfig.Redis.Setup)

	}
}
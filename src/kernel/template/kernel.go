package template

import (
	"log"
	"kernel/public"
	"ChunkLib/fileSystem"
	"ChunkLib/template"
)

func Init(){
	// 缓存模板 - 启动立即进行缓存
	if !public.WebSiteConfig.IsAllStatic {
		log.Println("初始化 [ 模板缓存 ] ...")

		basePath,err := fileSystem.GetMyPath()

		if err != nil {
			log.Fatal(err)
		}

		template.CacheHtmlTemplate(basePath+public.WebSiteConfig.TemplateUrl)
	}
}
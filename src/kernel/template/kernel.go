package template

import (
	"log"
	"kernel/public"
	"utils/ChunkLib/fileSystem"
	"utils/ChunkLib/template"
)

func Init(){
	/*
		缓存模板 - 启动立即进行缓存
		*/



	if !public.WebSiteConfig.IsAllStatic {

		log.Print("初始化 [ 模板缓存 ] ... \n\n")

		basePath,err := fileSystem.GetMyPath()

		if err != nil {
			log.Fatal(err)
		}

		template.CacheHtmlTemplate(basePath+public.WebSiteConfig.TemplateUrl)
	}


}
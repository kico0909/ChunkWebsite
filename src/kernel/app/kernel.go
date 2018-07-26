package app

/*
服务器的内核 用于启动服务器
*/

import (
	"log"
	"net/http"
	"time"
	"kernel/public"
	"kernel/route"
	"golang.org/x/crypto/acme/autocert"
	"strings"
	"crypto/tls"
	"golang.org/x/net/http2"
	"strconv"
)

func ServerStart(){


	if !public.WebSiteConfig.TLS.Open{
		server := &http.Server{
			// 地址及端口号
			Addr: `:`+strconv.FormatInt(public.WebSiteConfig.WebPort, 10),

			// 读取超时时间
			ReadTimeout: 10 * time.Second,

			// 写入超时时间
			WriteTimeout: 10 * time.Second,

			// 头字节限制
			MaxHeaderBytes: 32<<20,

			// 配置路由
			Handler: &route.Interceptor,
		}

		log.Println("服务器启动完成 [ 端口:"+strconv.FormatInt(public.WebSiteConfig.WebPort, 10)+" ]...\n\n")

		log.Fatal(server.ListenAndServe())
	}else{
		https_domain := strings.Split(public.WebSiteConfig.TLS.Domain, ",")

		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(https_domain...), //your domain here
			Cache:      autocert.DirCache("certs"),     //folder for storing certificates
			Email:      public.WebSiteConfig.TLS.Email,
		}
		// 80端口 301 重定向
		go http.ListenAndServe(":http", certManager.HTTPHandler(nil)) // 支持 http-01

		// server 配置
		server := &http.Server{
			Addr: ":https",
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
				NextProtos:     []string{http2.NextProtoTLS, "http/1.1"},
				MinVersion:     tls.VersionTLS12,
			},
			MaxHeaderBytes: 32<<20,
			Handler: &route.Interceptor,
		}

		log.Print(" TLS 服务器启动完成 ...\n\n")
		log.Fatal(server.ListenAndServeTLS("", ""))

	}
}


package app

/*
服务器的内核 用于启动服务器
*/

import (
	"log"
	"net/http"
	"time"
	"kernel/public"
	"strconv"
	"strings"
	"golang.org/x/crypto/acme/autocert"
	"crypto/tls"
	"golang.org/x/net/http2"
	"ChunkLib/cas"
	"github.com/gorilla/mux"
)


type RouterType interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func ServerStart(router RouterType){
	// 不启用HTTPS
	if !public.WebSiteConfig.TLS.Open{
		normalServerStart(router)
	}
	// 启用 HTTPS 并且 自动申请使用和续期let's Encrypt证书
	if public.WebSiteConfig.TLS.LetsEncrypt {
		httpsLetsServerStart(router)
	}
	httpsNormalServerStart(router)
}


// 非HTTPS服务器
func normalServerStart (router RouterType) {

	server := &http.Server{
		// 地址及端口号
		Addr: `:`+strconv.FormatInt(public.WebSiteConfig.WebPort, 10),

		// 读取超时时间
		ReadTimeout: 10 * time.Second,

		// 写入超时时间
		WriteTimeout: 10 * time.Second,

		// 头字节限制
		MaxHeaderBytes: 32<<20,

	}
	if router == nil {
		log.Fatalln("router为空!!!")

	}

	if public.WebSiteConfig.Cas.Key {

		server.Handler = router

	}else{
		server.Handler= router
	}

	log.Println("服务器启动完成 [ 端口:"+strconv.FormatInt(public.WebSiteConfig.WebPort, 10)+" ]...\n\n")

	log.Fatal(server.ListenAndServe())


}

// 启动自动申请let's encrypt 证书的服务器
func httpsLetsServerStart(router RouterType){
	https_domain := strings.Split(public.WebSiteConfig.TLS.LetsEncryptOpt.Domain, ",")

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(https_domain...), //your domain here
		Cache:      autocert.DirCache("certs"),     //folder for storing certificates
		Email:      public.WebSiteConfig.TLS.LetsEncryptOpt.Email,
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
	}

	if public.WebSiteConfig.Cas.Key {
		server.Handler = router

	}else{
		server.Handler= router
	}


	log.Print(" TLS 服务器启动完成 ...\n\n")
	log.Fatal(server.ListenAndServeTLS("", ""))
}

// 启动https服务器,需要填写证书路径
func httpsNormalServerStart(router interface{}){
	// 启用 HTTPS 直接加载证书
	server := &http.Server{

		// 地址及端口号
		Addr: `:`+strconv.FormatInt(public.WebSiteConfig.WebPort, 10),

		// 读取超时时间
		ReadTimeout: 10 * time.Second,

		// 写入超时时间
		WriteTimeout: 10 * time.Second,

		// 头字节限制
		MaxHeaderBytes: 32<<20,
	}

	if public.WebSiteConfig.Cas.Key {
		server.Handler = router.(*cas.CasRouter)

	}else{
		server.Handler= router.(*mux.Router)
	}

	log.Println("HTTPS 服务器启动完成 [ 端口:"+strconv.FormatInt(public.WebSiteConfig.WebPort, 10)+" ]...\n\n")

	log.Fatal(server.ListenAndServeTLS(public.WebSiteConfig.TLS.CertPath, public.WebSiteConfig.TLS.KeyPath))
}


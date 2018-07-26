package command

import (
	"kernel/public"
	"kernel/config"
	"kernel/route"
	"kernel/template"
	"kernel/session"
	"kernel/mysql"
	"kernel/redis"
	"kernel/cas"
	"kernel/app"
	"utils/ChunkLib/fileSystem"
	"flag"
	"os/exec"
	"os"
	"log"
	"strconv"
	"time"
)

const infoPath string = "./pid.txt"

var deamon = flag.Bool("d", false, "服务器静默运行模式!")
var Comm = flag.String("c","stop", "服务器执行的操作[ start:启动 | stop:停止 ]")


// 服务器参数处理
func ArgmentsHandler(){

	switch *Comm {

	case "start":
		saveStartInfos(strconv.FormatInt(int64(os.Getpid()), 10))
		serverStart()
		break

	case "stop":
		serverStop(loadStartInfos())
		break

	case "restart":
		serverStop(loadStartInfos())
		time.Sleep(2 * time.Second)
		saveStartInfos(strconv.FormatInt(int64(os.Getpid()), 10))
		serverStart()
		break

	default:

	}
}

// 服务器初始化与启动
func serverStart(){

	// 读取配置文件
	public.WebSiteConfig = config.Init()

	// 初始化配置路由
	route.Init()

	// 初始化模板
	template.Init()

	// 初始化session
	session.Init()

	// 初始化mysql
	mysql.Init()

	// 初始化redis
	redis.Init()

	// 初始化cas
	cas.Init()

	// 服务器启动
	app.ServerStart()	// 服务启动

}

// 服务器停止
func serverStop(pid string){

	var as []string = []string{"-9",pid}

	cmd := exec.Command("kill", as...)

	if cmd.Start() != nil {
		log.Fatal("关闭服务执行失败!")
	}else{
		log.Println("pid:[ " + pid + " ]进程被移除")
	}

	cmd = exec.Command("rm", "-rf", infoPath)
	cmd.Start()

}

// 记录启动信息PID文件
func saveStartInfos(pid string){
	f, err := os.OpenFile(infoPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, errs := f.WriteString(pid)
	if errs != nil {
		log.Fatal(errs)
	}
}

// 读取启动PID信息文件
func loadStartInfos()string{

	cont, err := fileSystem.ReadFile(infoPath)

	if err != nil {
		log.Fatal("PID记录文件无法读取, 请手动结束应用!")
	}

	return string(cont)
}

func init () {

	if !flag.Parsed() {
		flag.Parse()
	}

	if *deamon && (*Comm == "start" || *Comm == "restart") {
		cmd := exec.Command(os.Args[0], "-c", *Comm)
		cmd.Start()
		*deamon = false
		os.Exit(0)
	}

}










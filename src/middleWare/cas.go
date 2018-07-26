package middleWare

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"github.com/astaxie/beego/session"
	"time"
	"kernel/public"
)

type CasReqReturn struct {
	ServiceResponse serviceResponse	`json:"serviceResponse"`
	Timeout int64	`json:"timeout"`
}
type serviceResponse struct{
	AuthenticationFailure map[string]interface{}	`json:"authenticationFailure"`
	AuthenticationSuccess map[string]interface{}	`json:"authenticationSuccess"`

}


type casHandler struct {
	verifyUrl string
	sessionInfoTimeout int64
	CasSessionInfoName string
}

var Cas casHandler

// 通过CAS 去登录
func (_self *casHandler) Login(w http.ResponseWriter, r *http.Request) {
	ss, _ := public.Session.SessionStart(w,r)
	if ss.Delete(_self.CasSessionInfoName) == nil {
		http.Redirect(w,r,_self.verifyUrl+"?service=" + getFullUrl(r), http.StatusFound)
	}

}

// 验证是否登录 1. 判断session 是否有用户信息, 且用户信息未过期, 2. 是否有TICKET
func (_self *casHandler) IsLogined(w http.ResponseWriter, r *http.Request)bool{
	ss, _ := public.Session.SessionStart(w,r)
	uf := ss.Get(_self.CasSessionInfoName)
	ticket := r.FormValue("ticket")

	// 没有session  和 ticket ==> 重新登录
	if uf == nil && len(ticket)<1 {
		_self.Login(w, r)
		return false
	}
	// 没有SESSION 或 session 已经过期 但是有ticket ===> 请求用户信息 -> 就刷新了session 有效时间
	if (uf == nil || (uf.(CasReqReturn)).Timeout<=time.Now().Unix()) && len(ticket)>0 {
		return _self.getUserInfo(w, r, ss)
	}
	// 有session 但已经过期 ===> 重新登录
	if (uf.(CasReqReturn)).Timeout<=time.Now().Unix(){
		_self.Login(w, r)
		return false
	}

	return true

}
// 验证是否登录
func (_self *casHandler) Logout(w http.ResponseWriter, r *http.Request){
	ss, _ := public.Session.SessionStart(w,r)
	if ss.Delete(_self.CasSessionInfoName) == nil {
		http.Redirect(w,r,_self.verifyUrl+"/logout?service="+getFullUrl(r), http.StatusFound)
	}
}


// 获得当前页面的路径
func getFullUrl(r *http.Request)string{
	var part1 string
	if r.TLS != nil {
		part1 = "https://"
	}else{
		part1 = "http://"
	}
	host := url.QueryEscape(part1 + r.Host)
	return host
}


// 设置验证地址
func (_self *casHandler)Set (url string, sessionTimeout int64, sessionName string) bool {
	_self.verifyUrl = url
	_self.sessionInfoTimeout = sessionTimeout
	_self.CasSessionInfoName = sessionName
	return true
}

// 通过TICKET获得用户验证信息 , 并保存至用户的 session 中
func (_self *casHandler) getUserInfo(w http.ResponseWriter, r *http.Request, ss session.Store)bool{

	var ret CasReqReturn
	url := _self.verifyUrl + "/serviceValidate?ticket=" + r.FormValue("ticket") + "&service=" + getFullUrl(r) + "&format=JSON"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return false
	}
	reqRes, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return false
	}
	defer reqRes.Body.Close()
	body, err := ioutil.ReadAll(reqRes.Body)

	json.Unmarshal(body, &ret)

	if ret.ServiceResponse.AuthenticationFailure != nil {	// 当获得数据失败
		_self.Login(w, r)
		return false
	}

	ret.Timeout = time.Now().Unix() + _self.sessionInfoTimeout
	ss.Set(_self.CasSessionInfoName,ret)

	return true
}



package middleWare

/*
	验证 SIGN 的中间件
*/

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"crypto/sha1"
	"strconv"
	"net/http"
	"kernel/public"
)

type appInfoHandlerType struct {
	ts string
	nonceStr string	// 随机字符串
}

var Sign appInfoHandlerType

func (_self *appInfoHandlerType) Init(){

}

func (_self *appInfoHandlerType) VerifySign(r *http.Request)bool{
	/*
			1. 读取到header的 sign, ts, nonceAStr, appkey
			2. 验证 appkey 是否有效, 有效则取出appsecret
			3. 根据appsecret 创建sign , 并和传入sign 对比
		*/

	sign := r.Header.Get("sign")
	ts := r.Header.Get("ts")
	nonceStr := r.Header.Get("nonceStr")
	appkey := r.Header.Get("appkey")

	mySign, err:= MakeSign(appkey, nonceStr, ts)

	if err!= nil ||  !(sign == mySign){
		return false
	}else {
		return true
	}
}

func (_self *appInfoHandlerType) CreateAppInfos(){

}

func (_self *appInfoHandlerType) CreateAppkey(){

}

func (_self *appInfoHandlerType) CreateSecret(){

}

func (_self *appInfoHandlerType) MakeSign(){

}

func init(){

}

// 字符串转MD5
func str2md5(str string)string{
	tmp := md5.New()
	tmp.Write([]byte(str))
	MD5Str := tmp.Sum(nil)
	return hex.EncodeToString(MD5Str)
}

// 字符串SHA-1
func str2sha1( str string)string{
	t := sha1.New()
	t.Write([]byte(str))
	sha1Str := t.Sum(nil)
	return hex.EncodeToString(sha1Str)
}

// 创建APPkey,APPsecret 基础子串
func createBaseCode(md5Code, timeStr string) string {
	return md5Code+timeStr
}

//// 子串随机
//func randStrs(baseCode string) string {
//	finalCode := ""
//	arrayCode := strings.Split(baseCode, "")
//
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	for _, i := range r.Perm(len(arrayCode)) {
//		finalCode += arrayCode[i]
//	}
//
//	return finalCode
//}

func findSecret (appKey string) (secret string, err error){
	sqlRes, err := public.Mysql.Query("select appsecret from app_info where appkey='"+appKey+"'")
	if err != nil {
		return "",err
	}
	tmp := sqlRes[0]
	//if !ok {
	//	return "", errors.New("none data found")
	//}
	appsecret, _ := tmp["appsecret"]
	return appsecret, nil
}



// sign 计算方式
/*
	1. 拼接字符串: nonceStr , appkey , unix时间戳(毫秒级)
	2. sha1加密字符串
*/
func MakeSign(appkey, noncestr, ts string)(string, error){
	appsecret,_ := findSecret(appkey)
	tmpStr := noncestr + "," + appsecret + "," + ts
	return str2sha1(tmpStr), nil

}

func makeAppKey(baseCode string) string{
	return str2md5(baseCode)
}

func makeAppSecret(baseCode string) string{
	return str2sha1(baseCode)
}

// 获得APPkey,APPsecret
func InitAppCodes()(appkey, appsecret string){

	timeStampNano := strconv.FormatInt(time.Now().UnixNano(), 10)
	baseCode := createBaseCode(str2md5(timeStampNano), timeStampNano)
	//makeAppKey(baseCode)

	appkey = makeAppKey(baseCode)
	appsecret = makeAppSecret(baseCode)
	return appkey, appsecret
}

// 头信息验证的有效性
func CheckReqEffectiveness(r *http.Request)bool{
	/*
		1. 读取到header的 sign, ts, nonceAStr, appkey
		2. 验证 appkey 是否有效, 有效则取出appsecret
		3. 根据appsecret 创建sign , 并和传入sign 对比
	*/

	sign := r.Header.Get("sign")
	ts := r.Header.Get("ts")
	nonceStr := r.Header.Get("nonceStr")
	appkey := r.Header.Get("appkey")

	mySign, err:= MakeSign(appkey, nonceStr, ts)

	if err!= nil ||  !(sign == mySign){
		return false
	}else {
		return true
	}

}
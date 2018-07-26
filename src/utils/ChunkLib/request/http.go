package request

import (
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"log"
)

// Post请求
func Post(url string, data string, header map[string]string )(string, error){

	return sendHttpRequest(url, "POST", data, header )

	return data, nil
}

// Get请求
func Get(url string, data map[string]interface{}, header map[string]string )(string, error){
	return "", nil
}


// 发起http 请求
func sendHttpRequest(url string, method string, data string, header map[string]string )(string, error){

	client := &http.Client{}

	request, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(data))

	if err != nil {
		log.Println(err)
		return "", err
	}

	//// 定义默认 Content-Type 如果没有传入定义
	//if _,ok := header["Content-Type"]; !ok {
	//	header["Content-Type"] = "application/x-www-form-urlencoded"
	//}
	for k,v := range header {
		request.Header.Set(k,v)
	}

	reqRes, err := client.Do(request)

	if err != nil {
		log.Println(err)
		return "", err
	}

	defer reqRes.Body.Close()

	body, err := ioutil.ReadAll(reqRes.Body)

	if err != nil {
		log.Println(err)
		return "",err
	}


	return string(body), nil
}


// request 请求的结果合成
type REeqResult struct {
	Code int	`json:"code"`	// 服务器请求状态
	Success bool `json:"success"`	// 请求结果
	Error string	`json:"message"`	// 错误信息
	Data interface{}	`json:"data"`	// 返回数据
}

func ReqResultCreat(success bool, err error, data interface{})string{

	error := ""
	code := 200

	if err != nil {
		error = err.Error()
		code = 400
	}else{
		error = "null"
	}

	if !success {
		code = 400
	}

	r := REeqResult{code,success, error, data}

	res, errorInfo := json.Marshal(r)

	if errorInfo != nil {
		return errorInfo.Error()
	}

	return string(res)

}

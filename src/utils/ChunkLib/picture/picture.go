package picture

import (
"io/ioutil"
"strconv"
"strings"
"encoding/base64"
"errors"
"time"
)

/*
图片转换功能
*/

type picture struct {

}

var imageUploadPath = "/static/upload/picture/"

var Pic picture

/*
base64字符串 转 图片
*/
func( _self *picture ) Base64ToPicture(base64String string) (string, error){

	// 不是 BASE64 字符串
	if len(base64String) < 20 {
		return "", errors.New("param not base64")
	}

	// 获得图片名
	name := strconv.FormatInt(time.Now().UnixNano(), 10)

	// 拆分处理 base64
	slice := strings.Split(base64String, "base64,")

	b, err := base64.StdEncoding.DecodeString(slice[1])

	if err != nil {
		return "",  errors.New("string slice false")
	}

	fileName := "." + imageUploadPath+name+".png"

	if ioutil.WriteFile(fileName, b, 0666) == nil {
		return ""+imageUploadPath+name+".png", nil
	}

	return "", errors.New("write file false")

}

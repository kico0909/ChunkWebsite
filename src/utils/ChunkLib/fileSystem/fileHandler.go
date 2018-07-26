/*
操作文件系统
*/
package fileSystem

import (
	"runtime"
	"os"
	"path/filepath"
	"errors"
	"io/ioutil"
	)

/*
获得当前的系统绝对路径
*/
func GetMyPath() (string,error) {

	switch runtime.GOOS {


	case "darwin":
		return os.Getwd()
		break;

	case "windows":
		return os.Getwd()
		break;

	case "linux":
		return filepath.Abs(filepath.Dir(os.Args[0]))
		break;


	}
	return "",errors.New("unkonw OS System ! ["+runtime.GOOS+"]")
}


/*
读取指定文件
*/
func ReadFile(filePath string)([]byte, error){
	return ioutil.ReadFile(filePath)
}

/*
检测创建指定文件
*/



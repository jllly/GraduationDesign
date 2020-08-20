package tool

import (
	"Homework_system/global"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadConfigFile(path string) {
	var (
		file        *os.File
		err         error
		filecontent []byte
	)

	file, err = os.Open(path)
	if err != nil {
		fmt.Println("加载配置文件失败")
		return
	}
	filecontent, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("加载配置文件失败")
		return
	}
	global.Gconfig = global.Config{}
	err = json.Unmarshal(filecontent, &global.Gconfig)
	if err != nil {
		fmt.Println("解析配置文件失败")
		return
	}
}

/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:28:49
 * @Description:
 */
package main

import (
	"fmt"
	"os"
	"runtime"

	"com.fanyunai.sqlHelper/source"
	"com.fanyunai.sqlHelper/source/config"
	"gopkg.in/yaml.v2"
)

func init() {
	by, err := os.ReadFile("config.yml")
	if err != nil {
		fmt.Println("读取文件失败")
		os.Exit(1)
	}
	err = yaml.Unmarshal(by, &config.Conf)
	if err != nil {
		fmt.Println("解析文件错误")
		os.Exit(1)
	}
	sysPath := "/"
	if runtime.GOOS == "windows" {
		sysPath = "\\"
	}
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前路径无效")
		os.Exit(1)
	}
	// data 觉得路径
	dataName := "data.yml"
	config.DataFilePath = path + sysPath + dataName
	// 解析文件
	by, err = os.ReadFile(dataName)
	if err != nil {
		fmt.Println("读取文件失败")
		os.Exit(1)
	}
	err = yaml.Unmarshal(by, &config.DataBase)
	if err != nil {
		fmt.Println("解析文件错误")
		os.Exit(1)
	}
}

func main() {
	// 创建路由
	source.CreateRouter()
}

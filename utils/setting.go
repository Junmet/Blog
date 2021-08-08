package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string

	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
)

// 一个包初始化的接口
func init ()  {
	file, err := ini.Load("config/config.ini") // Load方法会返回一个file指针，和一个error （*File,error）
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:",err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File)  {  //把file传进去，接收的是*ini.File指针
	AppMode =file.Section("server").Key("AppMode").MustString("debug")
	HttpPort =file.Section("HttpPort").Key("HttpPort").MustString(":3000")
}
func LoadData(file *ini.File)  {
	Db =file.Section("database").Key("Db").MustString("mysql")
	DbHost =file.Section("database").Key("DbHost").MustString("localhost")
	DbPort =file.Section("database").Key("DbPort").MustString("3306")
	DbUser =file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord =file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName =file.Section("database").Key("DbName").MustString("ginblog")
}
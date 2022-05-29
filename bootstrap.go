package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Server struct {
	AppMode  string
	HTTPPort string
}

// InitServer 初始化服务变量
func InitServer() *Server {
	var server Server
	if _, err := toml.DecodeFile("./conf/app.toml", &server); err != nil {
		// 打日志
		fmt.Println(err)
	}
	return &server
}

func InitLogger() error {
	return nil
}

func initMySQL() {

}

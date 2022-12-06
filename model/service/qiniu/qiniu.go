package qiniu

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type QiniuConf struct {
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
}

var (
	qiniuConf QiniuConf // 数据库配置

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string

	err error // 全局错误
)

func init() {
	if _, err = toml.DecodeFile("./conf/qiniu.toml", &qiniuConf); err != nil {
		// TODO 打日志
		fmt.Println(err)
	}
	AccessKey = qiniuConf.AccessKey
	SecretKey = qiniuConf.SecretKey
	Bucket = qiniuConf.Bucket
	QiniuServer = qiniuConf.QiniuServer
}

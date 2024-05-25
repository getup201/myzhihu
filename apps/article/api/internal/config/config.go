package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	//继承
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	//阿里云OSS配置信息 结构体
	ArticleRPC zrpc.RpcClientConf
	UserRPC    zrpc.RpcClientConf

	Oss struct {
		Endpoint         string
		AccessKeyId      string
		AccessKeySecret  string
		BucketName       string
		ConnectTimeout   int64 `json:",optional"`
		ReadWriteTimeout int64 `json:",optional"`
	}
}

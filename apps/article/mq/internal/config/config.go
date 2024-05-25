package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	//点赞配置 和 文章变更配置
	KqConsumerConf        kq.KqConf
	ArticleKqConsumerConf kq.KqConf
	Datasource            string
	BizRedis              redis.RedisConf
	//es 的配置config  后面再配置
	Es struct {
		Addresses []string
		Username  string
		Password  string
	}
	UserRPC zrpc.RpcClientConf
}

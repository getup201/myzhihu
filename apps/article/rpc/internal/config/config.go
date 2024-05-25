package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	CacheRedis cache.CacheConf
	BizRedis   redis.RedisConf
	// CacheRedis 是gozero默认生成的缓存相关的配置
	// BizRedis 是跟自己的业务相关的缓存配置
	// 先不加 加
	// Consul consul.Conf
}

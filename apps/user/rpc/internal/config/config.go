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
	//BizRedis 是 Redis 的另一个实现，它主要用于实时数据处理和分析。
	//BizRedis 提供了一些高级功能，例如数据分片、数据复制、事务处理、发布/订阅模式等。
	//这些功能使得 BizRedis 可以处理大量的并发请求，并支持实时数据处理和分析。666
	BizRedis redis.RedisConf
	// 生成jwt时要用
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	// 后面再加这个
	// Consul     consul.Conf
}

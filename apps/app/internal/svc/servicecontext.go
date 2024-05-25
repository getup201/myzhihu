package svc

import (
	"myzhihu/apps/app/internal/config"
	"myzhihu/apps/user/rpc/user"
	"myzhihu/pkg/interceptors"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 定义一个userRPC 实例并注册自定义拦截器   客户端的拦截器 rpc错误码转自定义XCode 因为api不能识别rpc的错误码
	// 注意，在定义错误码的时候千万不能冲突，各业务服务一定要协商好，比如applet-api服务的错误码以1开头，用户服务的业务错误码以2开头。
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userRPC),
		BizRedis: redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}

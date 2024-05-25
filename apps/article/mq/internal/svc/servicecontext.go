package svc

import (
	"myzhihu/apps/article/mq/internal/config"
	"myzhihu/apps/article/mq/internal/model"
	"myzhihu/apps/user/rpc/user"
	"myzhihu/pkg/es"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

// 初始化一些依赖
type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel
	BizRedis     *redis.Redis
	UserRPC      user.User //聚合了一些用户的信息
	Es           *es.Es    //Es先不加
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}

	conn := sqlx.NewMysql(c.Datasource)
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(conn),
		BizRedis:     rds,
		UserRPC:      user.NewUser(zrpc.MustNewClient(c.UserRPC)),
		Es: es.MustNewEs(&es.Config{
			Addresses: c.Es.Addresses,
			Username:  c.Es.Username,
			Password:  c.Es.Password,
		}),
	}
}

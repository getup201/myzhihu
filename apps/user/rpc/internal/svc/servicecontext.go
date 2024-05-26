package svc

import (
	"myzhihu/apps/user/rpc/internal/config"
	"myzhihu/apps/user/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	// 加userauth
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config: c,
		//传入的参数主要是数据库和redis  (conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option)
		UserModel:     model.NewUserModel(conn, c.CacheRedis),
		UserAuthModel: model.NewUserAuthModel(conn, c.CacheRedis),
	}
}

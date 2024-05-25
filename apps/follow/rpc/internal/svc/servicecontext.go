package svc

import (
	"myzhihu/apps/follow/rpc/internal/config"
	"myzhihu/apps/follow/rpc/internal/model"
	"myzhihu/pkg/orm"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

// 由数据库orm的DB  和自定义的几个model（因为是用的gorm的DB操作 所有要自己定义）（之前是gozero自动生成的）
type ServiceContext struct {
	Config           config.Config
	DB               *orm.DB
	FollowModel      *model.FollowModel
	FollowCountModel *model.FollowCountModel
	BizRedis         *redis.Redis
}

// MustNewMysql MustNewRedis 的Must表示强依赖
func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})

	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})

	return &ServiceContext{
		Config:           c,
		DB:               db,
		FollowModel:      model.NewFollowModel(db.DB),
		FollowCountModel: model.NewFollowCountModel(db.DB),
		BizRedis:         rds,
	}
}

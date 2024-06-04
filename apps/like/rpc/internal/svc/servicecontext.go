package svc

import (
	"myzhihu/apps/like/rpc/internal/config"
	"myzhihu/apps/like/rpc/internal/model"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	BizRedis        *redis.Redis
	KqPusherClient  *kq.Pusher
	LikeRecordModel model.LikeRecordModel
	LikeCountModel  model.LikeCountModel
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
	return &ServiceContext{
		Config:          c,
		BizRedis:        rds,
		KqPusherClient:  kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		LikeRecordModel: model.NewLikeRecordModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		LikeCountModel:  model.NewLikeCountModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
	}
}

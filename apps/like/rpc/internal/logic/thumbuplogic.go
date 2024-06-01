package logic

import (
	"context"
	"encoding/json"

	"myzhihu/apps/like/rpc/internal/svc"
	"myzhihu/apps/like/rpc/internal/types"
	"myzhihu/apps/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type ThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送Kafka消息 发送Kafka消息
// rpc服务作为 kafka消息的生产端
func (l *ThumbupLogic) Thumbup(in *service.ThumbupRequest) (*service.ThumbupResponse, error) {
	// TODO 逻辑暂时忽略
	// 1. 查询是否点过赞
	// 2. 计算当前内容的总点赞数和点踩数
	// 点赞数据写入到点赞表中

	msg := &types.ThumbupMsg{
		BizId:    in.BizId,
		ObjId:    in.ObjId,
		UserId:   in.UserId,
		LikeType: in.LikeType,
	}
	//一开始消费端没反应是因为 rpc的yaml中的Topic没有配置对应
	//向发送kafka消息，异步  调了一个协程
	threading.GoSafe(func() {
		data, err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("[Thumbup] marshal msg: %v error: %v", msg, err)
			return
		}
		//因为在ServiceContext中定义了 kafka初始的client  KqPusherClient  就可以用Pus函数来生成消息
		err = l.svcCtx.KqPusherClient.Push(string(data))
		if err != nil {
			l.Logger.Errorf("[Thumbup] kq push data: %s error: %v", data, err)
		}

	})
	// fmt.Println("after 	threading.GoSafe(func()")
	return &service.ThumbupResponse{}, nil
}

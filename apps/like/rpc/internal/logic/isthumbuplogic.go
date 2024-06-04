package logic

import (
	"context"

	"myzhihu/apps/like/rpc/internal/model"
	"myzhihu/apps/like/rpc/internal/svc"
	"myzhihu/apps/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsThumbupLogic {
	return &IsThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsThumbupLogic) IsThumbup(in *service.IsThumbupRequest) (*service.IsThumbupResponse, error) {
	// todo: add your logic here and delete this line
	// 查询数据库 业务对象 用户 点赞对象
	isthumbupresponse := make(map[int64]*service.UserThumbup)
	likeRecord, err := l.svcCtx.LikeRecordModel.FindOneByBizIdObjIdUserId(l.ctx, in.BizId, in.TargetId, in.UserId)
	if err != nil && err != model.ErrNotFound {
		logx.Errorf("DB querry error: %v", err)
		return nil, err
	}
	if err == model.ErrNotFound {
		return &service.IsThumbupResponse{
			UserThumbups: isthumbupresponse,
		}, nil
	}
	// fmt.Println(in)
	// fmt.Println(likeRecord)
	//已经点赞的情况  isthumbupresponse的key就是用户ID
	isthumbupresponse[likeRecord.UserId] = &service.UserThumbup{
		UserId:      likeRecord.UserId,
		ThumbupTime: likeRecord.CreateTime.Unix(),
		LikeType:    int32(likeRecord.LikeType),
	}
	return &service.IsThumbupResponse{
		UserThumbups: isthumbupresponse,
	}, nil
}

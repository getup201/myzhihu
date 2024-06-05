package logic

import (
	"context"

	"myzhihu/apps/like/rpc/internal/code"
	"myzhihu/apps/like/rpc/internal/model"
	"myzhihu/apps/like/rpc/internal/svc"
	"myzhihu/apps/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelThumbupLogic {
	return &CancelThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 增加取消点赞功能
func (l *CancelThumbupLogic) CancelThumbup(in *service.CancelThumbupRequest) (*service.CancelThumbupResponse, error) {
	// todo: add your logic here and delete this line
	// 先判断是否点过赞
	isthumbupLogic := NewIsThumbupLogic(l.ctx, l.svcCtx)
	isthumbup, err := isthumbupLogic.IsThumbup(&service.IsThumbupRequest{BizId: in.BizId, TargetId: in.ObjId, UserId: in.UserId})
	if err != nil {
		return nil, err
	}
	//没点过赞 或者点赞类型 与请求类型不一样
	if len(isthumbup.UserThumbups) == 0 || isthumbup.UserThumbups[in.UserId].LikeType != in.LikeType {
		return nil, code.CancelThumbupFailed
	}
	//取消点赞逻辑  先更新likerecord表 先要把点赞记录找出来
	likeRecord, err := l.svcCtx.LikeRecordModel.FindOneByBizIdObjIdUserId(l.ctx, in.BizId, in.ObjId, in.UserId)
	if err != nil {
		l.Logger.Errorf("LikeRecord DB error: %v", err)
		return nil, err
	}
	err = l.svcCtx.LikeRecordModel.Update(l.ctx, &model.LikeRecord{
		Id:       likeRecord.Id,
		BizId:    in.BizId,
		ObjId:    in.ObjId,
		UserId:   in.UserId,
		LikeType: -1,
	})
	if err != nil {
		l.Logger.Errorf("LikeRecord Update req: %v error: %v", in, err)
		return nil, err
	}

	//再更新likeCount表  更新表时一定要带上主键ID
	likeCount, err := l.svcCtx.LikeCountModel.FindOneByBizIdObjId(l.ctx, in.BizId, in.ObjId)
	if err != nil {
		logx.Errorf("LikeCount DB querry error: %v", err)
		return nil, err
	}
	err = l.svcCtx.LikeCountModel.Update(l.ctx, &model.LikeCount{
		Id:         likeCount.Id,
		BizId:      in.BizId,
		ObjId:      in.ObjId,
		LikeNum:    likeCount.LikeNum - int64(in.LikeType),          //为零时不变 为一时减一
		DislikeNum: likeCount.DislikeNum - (1 - int64(in.LikeType)), //为零时减一  为一时不变
	})
	if err != nil {
		l.Logger.Errorf("LikeCount Update req: %v error: %v", in, err)
		return nil, err
	}

	return &service.CancelThumbupResponse{
		BizId:      in.BizId,
		ObjId:      in.ObjId,
		LikeNum:    likeCount.LikeNum - int64(in.LikeType),
		DislikeNum: likeCount.DislikeNum - (1 - int64(in.LikeType)),
	}, nil
}

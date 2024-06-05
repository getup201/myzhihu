package logic

import (
	"context"

	"myzhihu/apps/like/rpc/internal/code"
	"myzhihu/apps/like/rpc/internal/model"
	"myzhihu/apps/like/rpc/internal/svc"
	"myzhihu/apps/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
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
	// 1. 查询是否点过赞 先实现IsThumbup逻辑
	// 2. 计算当前内容的总点赞数和点踩数
	// 3. 将点赞数据写入到点赞表中
	// 2024-6-2待办 写点赞逻辑 和 es部分的学习  看一下后面的生产消息不改的话会不会生成两条消息
	if in.LikeType != 0 && in.LikeType != 1 {
		return nil, code.LikeTypeInvalid
	}

	isthumbupLogic := NewIsThumbupLogic(l.ctx, l.svcCtx)
	isthumbup, err := isthumbupLogic.IsThumbup(&service.IsThumbupRequest{BizId: in.BizId, TargetId: in.ObjId, UserId: in.UserId})
	if err != nil {
		return nil, err
	}
	//map为空 未点赞过 插入数据
	if len(isthumbup.UserThumbups) == 0 {
		_, err := l.svcCtx.LikeRecordModel.Insert(l.ctx, &model.LikeRecord{
			BizId:    in.BizId,
			ObjId:    in.ObjId,
			UserId:   in.UserId,
			LikeType: int64(in.LikeType),
		})
		if err != nil {
			l.Logger.Errorf("LikeRecord Insert req: %v error: %v", in, err)
			return nil, err
		}
		likeCount, err := l.svcCtx.LikeCountModel.FindOneByBizIdObjId(l.ctx, in.BizId, in.ObjId)
		//LikeType为1是点赞 LikeType为0是点踩
		//likeCount为空表示没人点赞过
		if err == model.ErrNotFound {
			l.svcCtx.LikeCountModel.Insert(l.ctx, &model.LikeCount{
				BizId:      in.BizId,
				ObjId:      in.ObjId,
				LikeNum:    int64(in.LikeType),
				DislikeNum: int64(1 - in.LikeType),
			})
			return &service.ThumbupResponse{
				BizId:      in.BizId,
				ObjId:      in.ObjId,
				LikeNum:    int64(in.LikeType),
				DislikeNum: int64(1 - in.LikeType),
			}, nil

		} else {
			//表示已有记录  别人点过赞
			err = l.svcCtx.LikeCountModel.Update(l.ctx, &model.LikeCount{
				Id:         likeCount.Id,
				BizId:      in.BizId,
				ObjId:      in.ObjId,
				LikeNum:    likeCount.LikeNum + int64(in.LikeType),
				DislikeNum: likeCount.DislikeNum + 1 - int64(in.LikeType),
			})
			if err != nil {
				l.Logger.Errorf("LikeCount Update req: %v error: %v", in, err)
				return nil, err
			}

			return &service.ThumbupResponse{
				BizId:      in.BizId,
				ObjId:      in.ObjId,
				LikeNum:    likeCount.LikeNum + int64(in.LikeType),
				DislikeNum: likeCount.DislikeNum + int64(1-in.LikeType),
			}, nil
		}

	} else {
		// 自己点过赞或者踩时
		// 点赞类型相同 不用操作 表示已经点过赞了
		// fmt.Println("自己点过赞或者踩时")
		// fmt.Println(isthumbup)
		if isthumbup.UserThumbups[in.UserId].LikeType == in.LikeType {
			likeCount, err := l.svcCtx.LikeCountModel.FindOneByBizIdObjId(l.ctx, in.BizId, in.ObjId)
			if err != nil {
				logx.Errorf("LikeCount DB querry error: %v", err)
				return nil, err
			}
			return &service.ThumbupResponse{
				BizId:      in.BizId,
				ObjId:      in.ObjId,
				LikeNum:    likeCount.LikeNum,
				DislikeNum: likeCount.DislikeNum,
			}, nil
		} else {
			//点赞类型不同时  先要在LikeRecord中把那条记录的主键找出来 再更新LikeRecord
			likeRecord, err := l.svcCtx.LikeRecordModel.FindOneByBizIdObjIdUserId(l.ctx, in.BizId, in.ObjId, in.UserId)
			if err != nil {
				l.Logger.Errorf("LikeRecord DB error: %v", err)
				return nil, err
			}
			//更新likerecord表   Update函数要把 BizId ObjId 这些带上才行
			err = l.svcCtx.LikeRecordModel.Update(l.ctx, &model.LikeRecord{
				Id:       likeRecord.Id,
				BizId:    in.BizId,
				ObjId:    in.ObjId,
				UserId:   in.UserId,
				LikeType: int64(in.LikeType),
			})
			if err != nil {
				l.Logger.Errorf("LikeRecord Update req: %v error: %v", in, err)
				return nil, err
			}

			likeCount, err := l.svcCtx.LikeCountModel.FindOneByBizIdObjId(l.ctx, in.BizId, in.ObjId)
			//正常来说应该不会出错  既然是点过赞的likecount中肯定会有记录
			if err != nil {
				logx.Errorf("LikeCount DB querry error: %v", err)
				return nil, err
			}

			// 分为两种情况 LikeType为-1（表示点过赞又取消了） 和 不为-1
			var LikeNumUpdate, DislikeNumUpdate int64
			if likeRecord.LikeType != -1 {
				LikeNumUpdate = 2*int64(in.LikeType) - 1 //让type 为零时减一 为一时加一
				DislikeNumUpdate = 1 - 2*int64(in.LikeType)
			} else {
				//点过赞又取消了时 只有加一的情况 不需要考虑减一
				LikeNumUpdate = int64(in.LikeType)
				DislikeNumUpdate = 1 - int64(in.LikeType)
			}
			//再更新LikeCount  更新表一定要带上主键ID
			err = l.svcCtx.LikeCountModel.Update(l.ctx, &model.LikeCount{
				Id:         likeCount.Id,
				BizId:      in.BizId,
				ObjId:      in.ObjId,
				LikeNum:    likeCount.LikeNum + LikeNumUpdate,
				DislikeNum: likeCount.DislikeNum + DislikeNumUpdate,
			})
			if err != nil {
				l.Logger.Errorf("LikeCount Update req: %v error: %v", in, err)
				return nil, err
			}
			return &service.ThumbupResponse{
				BizId:      in.BizId,
				ObjId:      in.ObjId,
				LikeNum:    likeCount.LikeNum + LikeNumUpdate,
				DislikeNum: likeCount.DislikeNum + DislikeNumUpdate,
			}, nil

		}
	}
}

// 投递消息到kafka部分 先把这部分删了
// msg := &types.ThumbupMsg{
// 	BizId:    in.BizId,
// 	ObjId:    in.ObjId,
// 	UserId:   in.UserId,
// 	LikeType: in.LikeType,
// }
// //一开始消费端没反应是因为 rpc的yaml中的Topic没有配置对应
// //向发送kafka消息，异步  调了一个协程
// threading.GoSafe(func() {
// 	data, err := json.Marshal(msg)
// 	if err != nil {
// 		l.Logger.Errorf("[Thumbup] marshal msg: %v error: %v", msg, err)
// 		return
// 	}
// 	//因为在ServiceContext中定义了 kafka初始的client  KqPusherClient  就可以用Pus函数来生成消息
// 	err = l.svcCtx.KqPusherClient.Push(string(data))
// 	if err != nil {
// 		l.Logger.Errorf("[Thumbup] kq push data: %s error: %v", data, err)
// 	}

// })

package logic

import (
	"context"
	"encoding/json"

	"myzhihu/apps/app/internal/svc"
	"myzhihu/apps/app/internal/types"
	"myzhihu/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//	type UserInfoResponse struct {
//		UserId   int64  `json:"user_id"`
//		Username string `json:"username"`
//		Avatar   string `json:"avatar"`  Avatar是头像
//	}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	//context中有Id 然后通过userpc 调用findbyID方法 从数据库中查询
	//首先要通过jwt 鉴权  在api调试工具中要 将Headers添加一个 Authorization 值为生成的token
	userId, err := l.ctx.Value(types.UserIdKey).(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if userId == 0 {
		return &types.UserInfoResponse{}, nil
	}
	u, err := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdRequest{
		UserId: userId,
	})
	if err != nil {
		logx.Errorf("FindById userId: %d error: %v", userId, err)
		return nil, err
	}

	return &types.UserInfoResponse{
		UserId:   u.UserId,
		Username: u.Username,
		Avatar:   u.Avatar,
	}, nil
}

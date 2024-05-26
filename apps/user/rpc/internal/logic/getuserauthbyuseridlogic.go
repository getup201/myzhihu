package logic

import (
	"context"

	"myzhihu/apps/user/rpc/internal/model"
	"myzhihu/apps/user/rpc/internal/svc"
	"myzhihu/apps/user/rpc/service"
	"myzhihu/apps/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *service.GetUserAuthByUserIdRequest) (*service.GetUserAuthyUserIdResponse, error) {
	// todo: add your logic here and delete this line
	// 调用数据库 跟AuthKey 一样
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByUserIdAuthType(l.ctx, in.UserId, in.AuthType)
	if err != nil && err != model.ErrNotFound {
		logx.Errorf("get user auth fail! error: %v", err)
		return nil, err
	}

	//copier.Copy  这个库是gorm作者的另一款新作 user对应looklook中的usercenter
	var respUserAuth user.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)

	//返回实际都一样
	return &service.GetUserAuthyUserIdResponse{
		UserAuth: &respUserAuth,
	}, nil
}

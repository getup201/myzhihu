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

type GetUserAuthByAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByAuthKeyLogic {
	return &GetUserAuthByAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 增加 getUserAuthByAuthKey 和 getUserAuthByUserId
func (l *GetUserAuthByAuthKeyLogic) GetUserAuthByAuthKey(in *service.GetUserAuthByAuthKeyRequest) (*service.GetUserAuthByAuthKeyResponse, error) {
	// todo: add your logic here and delete this line
	// 调用数据库
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByAuthTypeAuthKey(l.ctx, in.AuthType, in.AuthKey)
	if err != nil && err != model.ErrNotFound {
		logx.Errorf("get user auth  fail! error: %v", err)
		return nil, err
	}

	//copier.Copy  这个库是gorm作者的另一款新作 user对应looklook中的usercenter
	var respUserAuth user.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)

	return &service.GetUserAuthByAuthKeyResponse{
		UserAuth: &respUserAuth,
	}, nil
}

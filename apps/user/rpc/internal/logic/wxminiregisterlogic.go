package logic

import (
	"context"

	"myzhihu/apps/user/rpc/internal/model"
	"myzhihu/apps/user/rpc/internal/svc"
	"myzhihu/apps/user/rpc/service"
	"myzhihu/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type WxMiniRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWxMiniRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniRegisterLogic {
	return &WxMiniRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 微信小程序的注册rpc
func (l *WxMiniRegisterLogic) WxMiniRegister(in *service.WxMiniRegisterRequest) (*service.WxMiniRegisterResponse, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		logx.Errorf("mobile:%s,err:%v", in.Mobile, err)
		return nil, err
	}
	if user != nil {
		logx.Errorf("Register user exists mobile:%s,err:%v", in.Mobile, err)
		return nil, err
	}

	var userId int64
	//这种设计方式的好处是将事务管理逻辑与业务逻辑分离，提高代码的可维护性和可读性。
	//由于go-zero的事务要在model中才能使用，但是我在model中做了个处理，把它在model中暴露出来，就可以在logic中使用
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.User)
		user.Mobile = in.Mobile

		insertResult, err := l.svcCtx.UserModel.Insert(ctx, session, user)
		if err != nil {
			logx.Errorf("Register db user Insert err:%v,user:%+v", err, user)
			return err
		}
		//LastInsertId 方法返回一个 int64 类型的值，表示数据库中最近一次插入操作生成的自增 ID。
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			logx.Errorf("Register db user insertResult.LastInsertId err:%v,user:%+v", err, user)
			return err
		}
		userId = lastId

		userAuth := new(model.UserAuth)
		userAuth.UserId = lastId
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
			logx.Errorf("Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	// 生成token 这里用pkg中的token生成方法 参考login 的api
	//登录会返回一个jwt token
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.JwtAuth.AccessSecret,
		AccessExpire: l.svcCtx.Config.JwtAuth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": userId,
		},
	})
	if err != nil {
		return nil, err
	}

	return &service.WxMiniRegisterResponse{
		UserId:       userId,
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
	}, nil
}

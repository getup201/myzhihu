package logic

import (
	"context"
	"time"

	"myzhihu/apps/user/rpc/internal/code"
	"myzhihu/apps/user/rpc/internal/model"
	"myzhihu/apps/user/rpc/internal/svc"
	"myzhihu/apps/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	// 当注册名字为空的时候，返回业务自定义错误码
	if len(in.Username) == 0 {
		return nil, code.RegisterNameEmpty
	}

	//插入到数据库中  这个是gozero 自动生成的方法  这个版本是没加session的 不支持事务 参考looklook 格式如下
	// if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {

	// }); err != nil {
	// 	return nil, err
	// }
	var userId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		ret, err := l.svcCtx.UserModel.Insert(l.ctx, session, &model.User{
			Username:   in.Username,
			Mobile:     in.Mobile,
			Avatar:     in.Avatar,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
		if err != nil {
			logx.Errorf("Register req: %v error: %v", in, err)
			return err
		}
		//LastInsertId 方法返回一个 int64 类型的值，表示数据库中最近一次插入操作生成的自增 ID。
		lastId, err := ret.LastInsertId()
		userId = lastId
		if err != nil {
			logx.Errorf("LastInsertId error: %v", err)
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	//返回注册的Id
	return &service.RegisterResponse{UserId: userId}, nil
}

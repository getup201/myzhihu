package logic

import (
	"context"
	"errors"
	"myzhihu/apps/app/internal/code"
	"myzhihu/apps/app/internal/svc"
	"myzhihu/apps/app/internal/types"
	"myzhihu/apps/user/rpc/user"
	"myzhihu/pkg/encrypt"
	"myzhihu/pkg/jwt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	prefixActivation = "biz#activation#%s"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.RegisterResponse, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	}
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		return nil, code.RegisterPasswdEmpty
	} else {
		req.Password = encrypt.EncPassword(req.Password)
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.VerificationCodeEmpty
	}
	// fmt.Println("req.VerificationCode :", req.VerificationCode)
	err := checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil {
		logx.Errorf("checkVerificationCode error: %v", err)
		return nil, err
	}
	mobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("EncMobile mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u != nil && u.UserId > 0 {
		return nil, code.MobileHasRegistered
	}
	//这里的注册方法是uerrpc中实现的 这里还没有实现  所以一开始是没有将用户加到数据库中的
	//调用UserRPC   UserRPC的Register方法会返回 插入到数据库的Id
	regRet, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Username: req.Name,
		Mobile:   mobile,
	})
	if err != nil {
		logx.Errorf("Register error: %v", err)
		return nil, err
	}

	//生成jwt token
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": regRet.UserId,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}

	//删除所使用的验证码
	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis)

	//可以多一个RefreshAfter  这个好像没啥用
	return &types.RegisterResponse{
		UserId: regRet.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}

func checkVerificationCode(rds *redis.Redis, mobile, code string) error {
	cacheCode, err := getActivationCache(mobile, rds)
	if err != nil {
		return err
	}

	//下面两个错误信息是返回给服务端  api测试时不会返回这个结果而是统一的500  INTERNAL_ERROR
	//在 Go 语言中，使用 errors.New("verification code failed") 通常是在服务端代码内部创建了一个新的错误对象。
	//这个错误对象可以用于服务端的日志记录、条件判断或者其他错误处理逻辑。
	if cacheCode == "" {
		return errors.New("verification code expired")
	}
	if cacheCode != code {
		return errors.New("verification code failed")
		//这样是将错误信息发送给客户端  明天看一下这个怎么写
	}
	// fmt.Println("checkVerificationCode :", code)
	// fmt.Println("cacheCode :", cacheCode)
	return nil
}

// bera领水交互
// Quai社区
// xion日常  社区研究
//spaceandtime  答题

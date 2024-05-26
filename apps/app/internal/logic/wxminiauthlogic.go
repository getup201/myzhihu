package logic

import (
	"context"

	"myzhihu/apps/app/internal/svc"
	"myzhihu/apps/app/internal/types"
	"myzhihu/apps/user/rpc/user"
	"myzhihu/pkg/jwt"

	//因为不能调用rpc的internal的参数
	userModelvars "myzhihu/apps/app/internal/vars"

	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 这个相当于加了一个 与login平行的路由  // Wechat-Mini auth 小程序鉴权
func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMIniAuthRequest) (resp *types.WXMiniAuthResponse, err error) {
	// todo: add your logic here and delete this line
	//1、Wechat-Mini 初始化小程序对象
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.AppId,
		AppSecret: l.svcCtx.Config.WxMiniConf.Secret,
		Cache:     cache.NewMemory(),
	})
	// 获取授权信息
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		// 使用自定义错误
		logx.Errorf("发起授权请求失败 error: %v , code : %s  , authResult : %+v", err, req.Code, authResult)
		return nil, err
	}
	//2、Parsing WeChat-Mini return data
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		logx.Errorf("解析数据失败 req : %+v , err: %v , authResult:%+v ", req, err, authResult)
		return nil, err
	}

	//3、bind user or login.
	// 检查用户是否存在
	var userId int64
	rpcRsp, err := l.svcCtx.UserRPC.GetUserAuthByAuthKey(l.ctx, &user.GetUserAuthByAuthKeyRequest{
		AuthType: userModelvars.UserAuthTypeSmallWX,
		AuthKey:  authResult.OpenID,
	})
	if err != nil {
		logx.Errorf("rpc call userAuthByAuthKey err : %v , authResult : %+v", err, authResult)
		return nil, err
	}
	// 不存在时 注册登录
	if rpcRsp.UserAuth == nil || rpcRsp.UserAuth.Id == 0 {
		//bind user.

		//Wechat-Mini Decrypted data 能够从小程序解析到手机号码
		mobile := userData.PhoneNumber
		// 小程序和手机登录 调用微信小程序注册rpc
		registerRsp, err := l.svcCtx.UserRPC.WxMiniRegister(l.ctx, &user.WxMiniRegisterRequest{
			AuthKey:  authResult.OpenID,
			AuthType: userModelvars.UserAuthTypeSmallWX,
			Mobile:   mobile,
		})
		if err != nil {
			logx.Errorf("UsercenterRpc.Register err :%v, authResult : %+v", err, authResult)
			return nil, err
		}

		return &types.WXMiniAuthResponse{
			UserId: registerRsp.UserId,
			Token: types.Token{
				AccessToken:  registerRsp.AccessToken,
				AccessExpire: registerRsp.AccessExpire,
			},
		}, nil

	} else {
		// 存在时直接登录 直接生成一个jwt token
		userId = rpcRsp.UserAuth.UserId
		token, err := jwt.BuildTokens(jwt.TokenOptions{
			AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
			AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
			Fields: map[string]interface{}{
				"userId": userId,
			},
		})
		if err != nil {
			return nil, err
		}

		return &types.WXMiniAuthResponse{
			UserId: userId,
			Token: types.Token{
				AccessToken:  token.AccessToken,
				AccessExpire: token.AccessExpire,
			},
		}, nil
	}
}

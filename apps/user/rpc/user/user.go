// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"myzhihu/apps/user/rpc/service"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FindByIdRequest              = service.FindByIdRequest
	FindByIdResponse             = service.FindByIdResponse
	FindByMobileRequest          = service.FindByMobileRequest
	FindByMobileResponse         = service.FindByMobileResponse
	GetUserAuthByAuthKeyRequest  = service.GetUserAuthByAuthKeyRequest
	GetUserAuthByAuthKeyResponse = service.GetUserAuthByAuthKeyResponse
	GetUserAuthByUserIdRequest   = service.GetUserAuthByUserIdRequest
	GetUserAuthyUserIdResponse   = service.GetUserAuthyUserIdResponse
	RegisterRequest              = service.RegisterRequest
	RegisterResponse             = service.RegisterResponse
	SendSmsRequest               = service.SendSmsRequest
	SendSmsResponse              = service.SendSmsResponse
	UserAuth                     = service.UserAuth
	WxMiniRegisterRequest        = service.WxMiniRegisterRequest
	WxMiniRegisterResponse       = service.WxMiniRegisterResponse

	User interface {
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		WxMiniRegister(ctx context.Context, in *WxMiniRegisterRequest, opts ...grpc.CallOption) (*WxMiniRegisterResponse, error)
		FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error)
		FindByMobile(ctx context.Context, in *FindByMobileRequest, opts ...grpc.CallOption) (*FindByMobileResponse, error)
		SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error)
		// 增加 getUserAuthByAuthKey 和 getUserAuthByUserId
		GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyRequest, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResponse, error)
		GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdRequest, opts ...grpc.CallOption) (*GetUserAuthyUserIdResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) WxMiniRegister(ctx context.Context, in *WxMiniRegisterRequest, opts ...grpc.CallOption) (*WxMiniRegisterResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.WxMiniRegister(ctx, in, opts...)
}

func (m *defaultUser) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.FindById(ctx, in, opts...)
}

func (m *defaultUser) FindByMobile(ctx context.Context, in *FindByMobileRequest, opts ...grpc.CallOption) (*FindByMobileResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.FindByMobile(ctx, in, opts...)
}

func (m *defaultUser) SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.SendSms(ctx, in, opts...)
}

// 增加 getUserAuthByAuthKey 和 getUserAuthByUserId
func (m *defaultUser) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyRequest, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.GetUserAuthByAuthKey(ctx, in, opts...)
}

func (m *defaultUser) GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdRequest, opts ...grpc.CallOption) (*GetUserAuthyUserIdResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.GetUserAuthByUserId(ctx, in, opts...)
}

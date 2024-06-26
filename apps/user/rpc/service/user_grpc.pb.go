// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: user.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	User_Register_FullMethodName             = "/service.User/Register"
	User_WxMiniRegister_FullMethodName       = "/service.User/wxMiniRegister"
	User_FindById_FullMethodName             = "/service.User/FindById"
	User_FindByMobile_FullMethodName         = "/service.User/FindByMobile"
	User_SendSms_FullMethodName              = "/service.User/SendSms"
	User_GetUserAuthByAuthKey_FullMethodName = "/service.User/getUserAuthByAuthKey"
	User_GetUserAuthByUserId_FullMethodName  = "/service.User/getUserAuthByUserId"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	WxMiniRegister(ctx context.Context, in *WxMiniRegisterRequest, opts ...grpc.CallOption) (*WxMiniRegisterResponse, error)
	FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error)
	FindByMobile(ctx context.Context, in *FindByMobileRequest, opts ...grpc.CallOption) (*FindByMobileResponse, error)
	SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error)
	// 增加 getUserAuthByAuthKey 和 getUserAuthByUserId
	GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyRequest, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResponse, error)
	GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdRequest, opts ...grpc.CallOption) (*GetUserAuthyUserIdResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) WxMiniRegister(ctx context.Context, in *WxMiniRegisterRequest, opts ...grpc.CallOption) (*WxMiniRegisterResponse, error) {
	out := new(WxMiniRegisterResponse)
	err := c.cc.Invoke(ctx, User_WxMiniRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error) {
	out := new(FindByIdResponse)
	err := c.cc.Invoke(ctx, User_FindById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FindByMobile(ctx context.Context, in *FindByMobileRequest, opts ...grpc.CallOption) (*FindByMobileResponse, error) {
	out := new(FindByMobileResponse)
	err := c.cc.Invoke(ctx, User_FindByMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error) {
	out := new(SendSmsResponse)
	err := c.cc.Invoke(ctx, User_SendSms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyRequest, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResponse, error) {
	out := new(GetUserAuthByAuthKeyResponse)
	err := c.cc.Invoke(ctx, User_GetUserAuthByAuthKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdRequest, opts ...grpc.CallOption) (*GetUserAuthyUserIdResponse, error) {
	out := new(GetUserAuthyUserIdResponse)
	err := c.cc.Invoke(ctx, User_GetUserAuthByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	WxMiniRegister(context.Context, *WxMiniRegisterRequest) (*WxMiniRegisterResponse, error)
	FindById(context.Context, *FindByIdRequest) (*FindByIdResponse, error)
	FindByMobile(context.Context, *FindByMobileRequest) (*FindByMobileResponse, error)
	SendSms(context.Context, *SendSmsRequest) (*SendSmsResponse, error)
	// 增加 getUserAuthByAuthKey 和 getUserAuthByUserId
	GetUserAuthByAuthKey(context.Context, *GetUserAuthByAuthKeyRequest) (*GetUserAuthByAuthKeyResponse, error)
	GetUserAuthByUserId(context.Context, *GetUserAuthByUserIdRequest) (*GetUserAuthyUserIdResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) WxMiniRegister(context.Context, *WxMiniRegisterRequest) (*WxMiniRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxMiniRegister not implemented")
}
func (UnimplementedUserServer) FindById(context.Context, *FindByIdRequest) (*FindByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedUserServer) FindByMobile(context.Context, *FindByMobileRequest) (*FindByMobileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByMobile not implemented")
}
func (UnimplementedUserServer) SendSms(context.Context, *SendSmsRequest) (*SendSmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedUserServer) GetUserAuthByAuthKey(context.Context, *GetUserAuthByAuthKeyRequest) (*GetUserAuthByAuthKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAuthByAuthKey not implemented")
}
func (UnimplementedUserServer) GetUserAuthByUserId(context.Context, *GetUserAuthByUserIdRequest) (*GetUserAuthyUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAuthByUserId not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_WxMiniRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxMiniRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).WxMiniRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_WxMiniRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).WxMiniRegister(ctx, req.(*WxMiniRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindById(ctx, req.(*FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FindByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_FindByMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindByMobile(ctx, req.(*FindByMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SendSms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendSms(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserAuthByAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAuthByAuthKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserAuthByAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserAuthByAuthKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserAuthByAuthKey(ctx, req.(*GetUserAuthByAuthKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserAuthByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAuthByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserAuthByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserAuthByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserAuthByUserId(ctx, req.(*GetUserAuthByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "wxMiniRegister",
			Handler:    _User_WxMiniRegister_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _User_FindById_Handler,
		},
		{
			MethodName: "FindByMobile",
			Handler:    _User_FindByMobile_Handler,
		},
		{
			MethodName: "SendSms",
			Handler:    _User_SendSms_Handler,
		},
		{
			MethodName: "getUserAuthByAuthKey",
			Handler:    _User_GetUserAuthByAuthKey_Handler,
		},
		{
			MethodName: "getUserAuthByUserId",
			Handler:    _User_GetUserAuthByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

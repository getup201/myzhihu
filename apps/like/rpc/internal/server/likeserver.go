// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package server

import (
	"context"

	"myzhihu/apps/like/rpc/internal/logic"
	"myzhihu/apps/like/rpc/internal/svc"
	"myzhihu/apps/like/rpc/service"
)

type LikeServer struct {
	svcCtx *svc.ServiceContext
	service.UnimplementedLikeServer
}

func NewLikeServer(svcCtx *svc.ServiceContext) *LikeServer {
	return &LikeServer{
		svcCtx: svcCtx,
	}
}

func (s *LikeServer) Thumbup(ctx context.Context, in *service.ThumbupRequest) (*service.ThumbupResponse, error) {
	l := logic.NewThumbupLogic(ctx, s.svcCtx)
	return l.Thumbup(in)
}

func (s *LikeServer) IsThumbup(ctx context.Context, in *service.IsThumbupRequest) (*service.IsThumbupResponse, error) {
	l := logic.NewIsThumbupLogic(ctx, s.svcCtx)
	return l.IsThumbup(in)
}

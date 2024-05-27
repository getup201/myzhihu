package logic

import (
	"context"

	"myzhihu/apps/article/api/internal/svc"
	"myzhihu/apps/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlesLogic {
	return &ArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticlesLogic) Articles(req *types.ArticlesRequest) (resp *types.ArticlesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

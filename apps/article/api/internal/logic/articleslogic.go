package logic

import (
	"context"
	"encoding/json"

	"myzhihu/apps/article/api/internal/svc"
	"myzhihu/apps/article/api/internal/types"
	"myzhihu/apps/article/rpc/pb"
	"myzhihu/pkg/xcode"

	"github.com/jinzhu/copier"
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
	//直接从context中可以取出userID  通过jwt存的
	//l.ctx.Value("userId")是接口类型通过.(json.Number)指定为json.Number类型
	//再通过.Int64()转为int类型
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		logx.Errorf("l.ctx.Value error: %v", err)
		return nil, xcode.NoLogin
	}
	//直接调用 底层rpc
	ArticlesResp, err := l.svcCtx.ArticleRPC.Articles(l.ctx, &pb.ArticlesRequest{
		UserId:    userId,
		ArticleId: req.ArticleId,
		Cursor:    req.Cursor,
		PageSize:  req.PageSize,
		SortType:  req.SortType,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.ArticleRPC.Articles req: %v UserId : %d ArticleId: %d error: %v", req, userId, req.ArticleId, err)
		return nil, err
	}

	var articlesresp types.ArticlesResponse
	_ = copier.Copy(&articlesresp, ArticlesResp)

	return &articlesresp, nil
}

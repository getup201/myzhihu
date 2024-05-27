package logic

import (
	"context"
	"encoding/json"

	"myzhihu/apps/article/api/internal/svc"
	"myzhihu/apps/article/api/internal/types"
	"myzhihu/apps/article/rpc/pb"
	"myzhihu/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleRequest) (resp *types.DeleteArticleResponse, err error) {
	// todo: add your logic here and delete this line
	//删除文章api 输入也是一个文章ID 可以先对输入做一些限制

	//直接从context中可以取出userID  通过jwt存的
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		logx.Errorf("l.ctx.Value error: %v", err)
		return nil, xcode.NoLogin
	}

	_, err = l.svcCtx.ArticleRPC.ArticleDelete(l.ctx, &pb.ArticleDeleteRequest{
		UserId:    userId,
		ArticleId: req.ArticleId,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.ArticleRPC.ArticleDelete req: %v ArticleId: %d error: %v", req, req.ArticleId, err)
		return &types.DeleteArticleResponse{
			DeleteMessage: "删除失败",
		}, err
	}

	return &types.DeleteArticleResponse{
		DeleteMessage: "删除成功",
	}, nil
}

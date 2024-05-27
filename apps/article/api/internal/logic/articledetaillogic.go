package logic

import (
	"context"

	"myzhihu/apps/article/api/internal/code"
	"myzhihu/apps/article/api/internal/svc"
	"myzhihu/apps/article/api/internal/types"
	"myzhihu/apps/article/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleDetailLogic {
	return &ArticleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleDetailLogic) ArticleDetail(req *types.ArticleDetailRequest) (resp *types.ArticleDetailResponse, err error) {
	// todo: add your logic here and delete this line
	// 文章详情 API
	// 可以先对输入做一些限制

	//调用 底层rpc来实现
	DetailResp, err := l.svcCtx.ArticleRPC.ArticleDetail(l.ctx, &pb.ArticleDetailRequest{
		ArticleId: req.ArticleId,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.ArticleRPC.ArticleDetail req: %v ArticleId: %d error: %v", req, req.ArticleId, err)
		return nil, err
	}

	if DetailResp.Article == nil {
		return nil, code.ArticleNOtFound
	}
	// 这个错误的原因是 detailResp.Article 的类型是 *pb.ArticleItem（指针类型），
	// 结构体 types.ArticleItem 需要的是值类型 types.ArticleItem。这两个类型虽然结构相同，
	// 但由于它们属于不同的包（pb 和 types），Go 认为它们是不同的类型，不能直接互换。
	// 为了避免这个错误，需要手动将 *pb.ArticleItem 转换为 types.ArticleItem。
	// return &types.ArticleDetailResponse{
	// 	Article: types.ArticleItem{
	// 		Id:           detailResp.Article.Id,
	// 		Title:        detailResp.Article.Title,
	// 		Content:      detailResp.Article.Content,
	// 		Description:  detailResp.Article.Description,
	// 		Cover:        detailResp.Article.Cover,
	// 		CommentCount: detailResp.Article.CommentCount,
	// 		LikeCount:    detailResp.Article.LikeCount,
	// 		PublishTime:  detailResp.Article.PublishTime,
	// 		AuthorId:     detailResp.Article.AuthorId,
	// 	},
	// }, nil

	//这里有个小技巧，很多同学感觉rpc服务返回的字段跟api定义差不多，每次都要手动去复制很麻烦
	// 可以使用copier.Copy 进行快速转换，这个库是gorm作者的另一款新作
	var detailresp types.ArticleDetailResponse
	_ = copier.Copy(&detailresp, DetailResp)

	return &detailresp, nil

}

package logic

import (
	"context"
	"strconv"
	"time"

	"myzhihu/apps/article/rpc/internal/code"
	"myzhihu/apps/article/rpc/internal/model"
	"myzhihu/apps/article/rpc/internal/svc"
	"myzhihu/apps/article/rpc/internal/types"
	"myzhihu/apps/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 代码路径 ~/beyond/application/article/rpc/internal/logic/publishlogic.go:27
// 在Publish方法中代用model生成的Insert方法进行数据插入 直接向数据库插入一条数据
// 调用api时要用 json格式
func (l *PublishLogic) Publish(in *pb.PublishRequest) (*pb.PublishResponse, error) {
	//参数校验 过滤非法参数
	if in.UserId <= 0 {
		return nil, code.UserIdInvalid
	}
	if len(in.Title) == 0 {
		return nil, code.ArticleTitleCantEmpty
	}
	if len(in.Content) == 0 {
		return nil, code.ArticleContentCantEmpty
	}
	ret, err := l.svcCtx.ArticleModel.Insert(l.ctx, &model.Article{
		AuthorId:    in.UserId,
		Title:       in.Title,
		Content:     in.Content,
		Description: in.Description,
		Cover:       in.Cover,
		Status:      types.ArticleStatusVisible, // 正常逻辑不会这样写，这里为了演示方便  真实业务中应该是 刚插入时状态为未审核
		PublishTime: time.Now(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		// PublishTime必须指定不指定会默认为0000-00-00产生bug  只有CreateTime  UpdateTime在gozero中能够默认
		// https://blog.csdn.net/Deng_Xian_Sheng/article/details/123215405
	})
	if err != nil {
		// fmt.Println("插入数据库失败")
		//log是这里输出的
		l.Logger.Errorf("Publish Insert req: %v error: %v", in, err)
		return nil, err
	}

	articleId, err := ret.LastInsertId()
	if err != nil {
		l.Logger.Errorf("LastInsertId error: %v", err)
		return nil, err
	}

	var (
		articleIdStr   = strconv.FormatInt(articleId, 10)
		publishTimeKey = articlesKey(in.UserId, types.SortPublishTime)
		likeNumKey     = articlesKey(in.UserId, types.SortLikeCount)
	)
	//同步写缓存 保证及时性 且写之前要判断缓存存在才能够添加 不然就会导致缓存不一致
	// 这部分代码注释起来 可以模拟写缓存异常  判断mq补偿的效果
	b, _ := l.svcCtx.BizRedis.ExistsCtx(l.ctx, publishTimeKey)
	if b {
		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, publishTimeKey, time.Now().Unix(), articleIdStr)
		if err != nil {
			logx.Errorf("ZaddCtx req: %v error: %v", in, err)
		}
	}
	b, _ = l.svcCtx.BizRedis.ExistsCtx(l.ctx, likeNumKey)
	if b {
		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, likeNumKey, 0, articleIdStr)
		if err != nil {
			logx.Errorf("ZaddCtx req: %v error: %v", in, err)
		}
	}

	return &pb.PublishResponse{ArticleId: articleId}, nil
}

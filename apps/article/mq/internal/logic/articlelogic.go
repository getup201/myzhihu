package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"myzhihu/apps/article/mq/internal/svc"
	"myzhihu/apps/article/mq/internal/types"
	"myzhihu/apps/user/rpc/user"

	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 消费文章变更数据
func (l *ArticleLogic) Consume(_, val string) error {
	//消费时打印日志 canal监听 发送到kafka
	logx.Infof("Consume msg val: %s", val)
	// fmt.Println("----------------------Consume正在执行------------------")
	var msg *types.CanalArticleMsg
	err := json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}

	return l.articleOperate(msg)
}

// 核心方法  异步处理 对文章变更进行补偿操作 保证最终一致性  点赞可以参考这里
func (l *ArticleLogic) articleOperate(msg *types.CanalArticleMsg) error {
	// fmt.Println("调用了articleOperate函数")
	if len(msg.Data) == 0 {
		return nil
	}

	var esData []*types.ArticleEsMsg
	// var err error
	for _, d := range msg.Data {
		status, _ := strconv.Atoi(d.Status)
		likNum, _ := strconv.ParseInt(d.LikeNum, 10, 64)
		articleId, _ := strconv.ParseInt(d.ID, 10, 64)
		authorId, _ := strconv.ParseInt(d.AuthorId, 10, 64)

		var err error
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", d.PublishTime, time.Local)
		publishTimeKey := articlesKey(d.AuthorId, 0)
		likeNumKey := articlesKey(d.AuthorId, 1)

		//根据文章状态判断
		switch status {
		case types.ArticleStatusVisible:
			b, _ := l.svcCtx.BizRedis.ExistsCtx(l.ctx, publishTimeKey)
			if b {
				_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, publishTimeKey, t.Unix(), d.ID)
				if err != nil {
					l.Logger.Errorf("ZaddCtx key: %s req: %v error: %v", publishTimeKey, d, err)
				}
			}
			b, _ = l.svcCtx.BizRedis.ExistsCtx(l.ctx, likeNumKey)
			if b {
				_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, likeNumKey, likNum, d.ID)
				if err != nil {
					l.Logger.Errorf("ZaddCtx key: %s req: %v error: %v", likeNumKey, d, err)
				}
			}
		case types.ArticleStatusUserDelete:
			_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, publishTimeKey, d.ID)
			if err != nil {
				l.Logger.Errorf("ZremCtx key: %s req: %v error: %v", publishTimeKey, d, err)
			}
			_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, likeNumKey, d.ID)
			if err != nil {
				l.Logger.Errorf("ZremCtx key: %s req: %v error: %v", likeNumKey, d, err)
			}
		}

		// es部分 用作搜索功能 执行一些复杂的查询(组合搜索.因为不能所有的字段都加索引)  调用下面的es核心代码
		// 调用userRPC   在数据库中查询到用户信息(包括username) 添加到es中
		u, err := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdRequest{
			UserId: authorId,
		})
		if err != nil {
			// 查询用
			l.Logger.Errorf("FindById userId: %d error: %v", authorId, err)
			return err
		}

		//聚合文章数据和用户数据
		esData = append(esData, &types.ArticleEsMsg{
			ArticleId:   articleId,
			AuthorId:    authorId,
			AuthorName:  u.Username,
			Title:       d.Title,
			Content:     d.Content,
			Description: d.Description,
			Status:      status,
			LikeNum:     likNum,
		})
	}
	//调用BatchUpSertToEs 将esData 写入到es中 (UPSert方法)
	err := l.BatchUpSertToEs(l.ctx, esData)
	if err != nil {
		l.Logger.Errorf("BatchUpSertToEs data: %v error: %v", esData, err)
	}

	return err
}

// es核心代码(批量) 主要是调用NewBulkIndexer 这个方法
func (l *ArticleLogic) BatchUpSertToEs(ctx context.Context, data []*types.ArticleEsMsg) error {
	if len(data) == 0 {
		return nil
	}
	//初始化
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: l.svcCtx.Es.Client,
		Index:  "article-index",
	})
	if err != nil {
		return err
	}

	for _, d := range data {
		v, err := json.Marshal(d)
		if err != nil {
			return err
		}

		payload := fmt.Sprintf(`{"doc":%s,"doc_as_upsert":true}`, string(v))
		//循环往 bulk中添加数据 写入到es中
		err = bi.Add(ctx, esutil.BulkIndexerItem{
			Action:     "update",
			DocumentID: fmt.Sprintf("%d", d.ArticleId),
			Body:       strings.NewReader(payload),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
			},
		})
		if err != nil {
			return err
		}
	}

	return bi.Close(ctx)
}

func articlesKey(uid string, sortType int32) string {
	return fmt.Sprintf("biz#articles#%s#%d", uid, sortType)
}

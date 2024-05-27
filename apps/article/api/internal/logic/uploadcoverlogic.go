package logic

import (
	"context"
	"fmt"
	"myzhihu/apps/article/api/internal/code"
	"myzhihu/apps/article/api/internal/svc"
	"myzhihu/apps/article/api/internal/types"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 这个表示10MB

type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 上传封面核心逻辑
func (l *UploadCoverLogic) UploadCover(req *http.Request) (*types.UploadCoverResponse, error) {
	_ = req.ParseMultipartForm(maxFileSize)
	file, handler, err := req.FormFile("cover")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//获取阿里云bucket
	bucket, err := l.svcCtx.OssClient.Bucket(l.svcCtx.Config.Oss.BucketName)
	if err != nil {
		logx.Errorf("get bucket failed, err: %v", err)
		return nil, code.GetBucketErr
	}

	//定义文件名 上传文件
	objectKey := genFilename(handler.Filename)
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		logx.Errorf("put object failed, err: %v", err)
		return nil, code.PutBucketErr
	}
	//返回这个文件的域名 只是一个字符串 后面再改
	return &types.UploadCoverResponse{CoverUrl: genFileURL(objectKey)}, nil
}

func genFilename(filename string) string {
	return fmt.Sprintf("%d_%s", time.Now().UnixMilli(), filename)
}

func genFileURL(objectKey string) string {
	return fmt.Sprintf("https://myzhihu-article.oss-cn-guangzhou.aliyuncs.com/%s", objectKey)
}

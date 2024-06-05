package code

import (
	"myzhihu/pkg/xcode"
)

// 点赞rpc
var (
	LikeTypeInvalid     = xcode.New(70001, "点赞类型无效")    // 点赞类型无效
	CancelThumbupFailed = xcode.New(70002, "无点赞记录无法取消") // 取消点赞失败
)

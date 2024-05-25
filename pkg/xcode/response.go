package xcode

import (
	"net/http"

	"myzhihu/pkg/xcode/types"
)

//	type XCode interface {
//		Error() string
//		Code() int
//		Message() string
//		Details() []interface{}
//	}
//
// 返回的code 结构体

// 在这里 any 和 interface{} 是等效的
func ErrHandler(err error) (int, any) {
	//http的err转换为  自定义的XCode
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}

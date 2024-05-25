package handler

import (
	"net/http"

	"myzhihu/apps/app/internal/logic"
	"myzhihu/apps/app/internal/svc"
	"myzhihu/apps/app/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			//如果发生错误，通过httpx.ErrorCtx(r.Context(), w, err)返回
			//这个是gozero 默认生成的
			// 自己在main函数里改 ErrorCtx中的doHandleError输入参数会变  默认会是nil自己设置了就不是nil了
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"myzhihu/apps/app/internal/logic"
	"myzhihu/apps/app/internal/svc"
	"myzhihu/apps/app/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取验证码的逻辑代码
func VerificationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerificationRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVerificationLogic(r.Context(), svcCtx)
		resp, err := l.Verification(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

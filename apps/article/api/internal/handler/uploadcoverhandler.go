package handler

import (
	"net/http"

	"myzhihu/apps/article/api/internal/logic"
	"myzhihu/apps/article/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadCoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadCoverLogic(r.Context(), svcCtx)
		resp, err := l.UploadCover(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

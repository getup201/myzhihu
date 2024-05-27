package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myzhihu/apps/article/api/internal/logic"
	"myzhihu/apps/article/api/internal/svc"
	"myzhihu/apps/article/api/internal/types"
)

func ArticleDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleDetailLogic(r.Context(), svcCtx)
		resp, err := l.ArticleDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package site

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"site/internal/handler/public/internal/logic/site"
	"site/internal/handler/public/internal/svc"
	"site/internal/handler/public/internal/types"
)

// List sites with filters and pagination
func ListSitesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListSitesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := site.NewListSitesLogic(r.Context(), svcCtx)
		resp, err := l.ListSites(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

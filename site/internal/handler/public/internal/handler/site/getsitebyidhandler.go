package site

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"site/internal/handler/public/internal/logic/site"
	"site/internal/handler/public/internal/svc"
	"site/internal/handler/public/internal/types"
)

// Get site details by ID
func GetSiteByIDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSiteByIDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := site.NewGetSiteByIDLogic(r.Context(), svcCtx)
		resp, err := l.GetSiteByID(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

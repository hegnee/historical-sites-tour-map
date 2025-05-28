package site

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"site/internal/handler/public/internal/logic/site"
	"site/internal/handler/public/internal/svc"
	"site/internal/handler/public/internal/types"
)

// Get dynamic information for a site (opening hours, ticket prices)
func GetSiteDynamicInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSiteDynamicInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := site.NewGetSiteDynamicInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetSiteDynamicInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

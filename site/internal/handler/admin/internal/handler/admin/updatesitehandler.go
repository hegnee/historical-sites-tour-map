package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"site/internal/handler/admin/internal/logic/admin"
	"site/internal/handler/admin/internal/svc"
	"site/internal/handler/admin/internal/types"
)

// (Admin) Update an existing site
func UpdateSiteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSiteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewUpdateSiteLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSite(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

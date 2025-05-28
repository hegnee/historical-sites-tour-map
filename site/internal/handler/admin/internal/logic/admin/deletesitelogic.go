package admin

import (
	"context"

	"site/internal/handler/admin/internal/svc"
	"site/internal/handler/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// (Admin) Delete a site
func NewDeleteSiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSiteLogic {
	return &DeleteSiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSiteLogic) DeleteSite(req *types.GetSiteByIDReq) (resp *types.DeleteSiteResp, err error) {
	// todo: add your logic here and delete this line

	return
}

package admin

import (
	"context"

	"site/internal/handler/admin/internal/svc"
	"site/internal/handler/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// (Admin) Update an existing site
func NewUpdateSiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSiteLogic {
	return &UpdateSiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSiteLogic) UpdateSite(req *types.UpdateSiteReq) (resp *types.Site, err error) {
	// todo: add your logic here and delete this line

	return
}

package admin

import (
	"context"

	"site/internal/handler/admin/internal/svc"
	"site/internal/handler/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// (Admin) Create a new site
func NewCreateSiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSiteLogic {
	return &CreateSiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSiteLogic) CreateSite(req *types.CreateSiteReq) (resp *types.Site, err error) {
	// todo: add your logic here and delete this line

	return
}

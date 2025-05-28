package site

import (
	"context"

	"site/internal/handler/public/internal/svc"
	"site/internal/handler/public/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSitesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// List sites with filters and pagination
func NewListSitesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSitesLogic {
	return &ListSitesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSitesLogic) ListSites(req *types.ListSitesReq) (resp *types.ListSitesResp, err error) {
	// todo: add your logic here and delete this line

	return
}

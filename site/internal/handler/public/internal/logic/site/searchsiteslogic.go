package site

import (
	"context"

	"site/internal/handler/public/internal/svc"
	"site/internal/handler/public/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSitesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Search sites by keyword
func NewSearchSitesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSitesLogic {
	return &SearchSitesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSitesLogic) SearchSites(req *types.SearchSitesReq) (resp *types.SearchSitesResp, err error) {
	// todo: add your logic here and delete this line

	return
}

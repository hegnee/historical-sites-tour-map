package site

import (
	"context"

	"site/internal/handler/public/internal/svc"
	"site/internal/handler/public/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSiteByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get site details by ID
func NewGetSiteByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSiteByIDLogic {
	return &GetSiteByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSiteByIDLogic) GetSiteByID(req *types.GetSiteByIDReq) (resp *types.Site, err error) {
	// todo: add your logic here and delete this line

	return
}

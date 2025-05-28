package site

import (
	"context"

	"site/internal/handler/public/internal/svc"
	"site/internal/handler/public/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSiteMultimediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get multimedia resources for a site
func NewGetSiteMultimediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSiteMultimediaLogic {
	return &GetSiteMultimediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSiteMultimediaLogic) GetSiteMultimedia(req *types.GetSiteMultimediaReq) (resp *types.GetSiteMultimediaResp, err error) {
	// todo: add your logic here and delete this line

	return
}

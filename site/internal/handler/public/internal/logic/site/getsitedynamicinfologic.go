package site

import (
	"context"

	"site/internal/handler/public/internal/svc"
	"site/internal/handler/public/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSiteDynamicInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get dynamic information for a site (opening hours, ticket prices)
func NewGetSiteDynamicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSiteDynamicInfoLogic {
	return &GetSiteDynamicInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSiteDynamicInfoLogic) GetSiteDynamicInfo(req *types.GetSiteDynamicInfoReq) (resp *types.DynamicInfo, err error) {
	// todo: add your logic here and delete this line

	return
}

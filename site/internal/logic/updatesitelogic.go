package logic

import (
	"context"

	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSiteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSiteLogic {
	return &UpdateSiteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSiteLogic) UpdateSite(in *pb.UpdateSiteReq) (*pb.SiteInfo, error) {
	// todo: add your logic here and delete this line

	return &pb.SiteInfo{}, nil
}

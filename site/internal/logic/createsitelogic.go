package logic

import (
	"context"

	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSiteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSiteLogic {
	return &CreateSiteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Admin methods
func (l *CreateSiteLogic) CreateSite(in *pb.CreateSiteReq) (*pb.SiteInfo, error) {
	// todo: add your logic here and delete this line

	return &pb.SiteInfo{}, nil
}

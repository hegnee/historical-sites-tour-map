package logic

import (
	"context"

	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSiteByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSiteByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSiteByIDLogic {
	return &GetSiteByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSiteByIDLogic) GetSiteByID(in *pb.GetSiteByIDReq) (*pb.SiteInfo, error) {
	// todo: add your logic here and delete this line

	return &pb.SiteInfo{}, nil
}

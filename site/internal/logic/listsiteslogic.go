package logic

import (
	"context"

	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSitesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSitesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSitesLogic {
	return &ListSitesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListSitesLogic) ListSites(in *pb.ListSitesReq) (*pb.ListSitesResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListSitesResp{}, nil
}

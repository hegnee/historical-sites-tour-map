package logic

import (
	"context"

	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSitesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSitesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSitesLogic {
	return &SearchSitesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSitesLogic) SearchSites(in *pb.SearchSitesReq) (*pb.SearchSitesResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchSitesResp{}, nil
}

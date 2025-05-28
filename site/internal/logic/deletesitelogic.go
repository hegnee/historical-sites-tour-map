package logic

import (
	"context"

	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSiteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSiteLogic {
	return &DeleteSiteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSiteLogic) DeleteSite(in *pb.DeleteSiteReq) (*pb.DeleteSiteResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteSiteResp{}, nil
}

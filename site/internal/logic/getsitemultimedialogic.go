package logic

import (
	"context"

	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSiteMultimediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSiteMultimediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSiteMultimediaLogic {
	return &GetSiteMultimediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSiteMultimediaLogic) GetSiteMultimedia(in *pb.GetSiteMultimediaReq) (*pb.GetSiteMultimediaResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSiteMultimediaResp{}, nil
}

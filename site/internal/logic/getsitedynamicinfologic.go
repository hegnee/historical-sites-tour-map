package logic

import (
	"context"

	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSiteDynamicInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSiteDynamicInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSiteDynamicInfoLogic {
	return &GetSiteDynamicInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSiteDynamicInfoLogic) GetSiteDynamicInfo(in *pb.GetSiteDynamicInfoReq) (*pb.DynamicInfo, error) {
	// todo: add your logic here and delete this line

	return &pb.DynamicInfo{}, nil
}

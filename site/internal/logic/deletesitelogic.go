package logic

import (
	"context"
	"errors"

	"site/internal/model"
	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// 1. Input Validation
	if in.SiteId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "SiteID cannot be empty")
	}

	// 2. Check Existence (Optional but recommended for consistent errors)
	_, err := l.svcCtx.SiteModel.FindOne(l.ctx, in.SiteId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "Site not found with ID: %s", in.SiteId)
		}
		l.Errorf("Error checking site existence before delete. ID: %s, Error: %v", in.SiteId, err)
		return nil, status.Errorf(codes.Internal, "Error checking site existence: %v", err)
	}

	// 3. Delete Record
	err = l.svcCtx.SiteModel.Delete(l.ctx, in.SiteId)
	if err != nil {
		// Note: Some Delete implementations might return ErrNotFound if the record was already gone.
		// If FindOne passed, this shouldn't be ErrNotFound unless there's a race condition.
		// However, if FindOne was omitted, checking for ErrNotFound here would be important.
		l.Errorf("Failed to delete site from database. ID: %s, Error: %v", in.SiteId, err)
		return nil, status.Errorf(codes.Internal, "Failed to delete site: %v", err)
	}

	// 4. Construct Response
	response := &pb.DeleteSiteResp{
		Message: "Site deleted successfully",
	}

	// 5. Return Response
	return response, nil
}

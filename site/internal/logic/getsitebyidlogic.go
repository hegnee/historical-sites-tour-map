package logic

import (
	"context"
	"database/sql" // For sql.Null types
	"errors"       // For errors.Is

	"site/internal/model" // Assuming model.ErrNotFound is defined here
	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if in.SiteId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "SiteID cannot be empty")
	}

	dbSite, err := l.svcCtx.SiteModel.FindOne(l.ctx, in.SiteId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "Site not found with ID: %s", in.SiteId)
		}
		l.Errorf("Failed to get site by ID %s: %v", in.SiteId, err)
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	// Map dbSite to pb.SiteInfo
	siteInfo := &pb.SiteInfo{
		Id:   dbSite.Id,
		Name: dbSite.Name,
		// CreatedAt and UpdatedAt are assumed to be non-nullable time.Time in the model
		// If they were sql.NullTime, similar .Valid checks and .Time access would be needed.
		CreatedAt: dbSite.CreatedAt.Unix(),
		UpdatedAt: dbSite.UpdatedAt.Unix(),
	}

	// Handle nullable fields by checking .Valid
	if dbSite.Description.Valid {
		siteInfo.Description = dbSite.Description.String
	} else {
		siteInfo.Description = "" // Default for optional string
	}

	if dbSite.Category.Valid {
		siteInfo.Category = dbSite.Category.String
	} else {
		siteInfo.Category = "" // Default for optional string
	}

	if dbSite.Region.Valid {
		siteInfo.Region = dbSite.Region.String
	} else {
		siteInfo.Region = "" // Default for optional string
	}

	if dbSite.HistoricalPeriod.Valid {
		siteInfo.HistoricalPeriod = dbSite.HistoricalPeriod.String
	} else {
		siteInfo.HistoricalPeriod = "" // Default for optional string
	}

	if dbSite.Latitude.Valid {
		siteInfo.Latitude = dbSite.Latitude.Float64
	} else {
		siteInfo.Latitude = 0.0 // Default for optional float
	}

	if dbSite.Longitude.Valid {
		siteInfo.Longitude = dbSite.Longitude.Float64
	} else {
		siteInfo.Longitude = 0.0 // Default for optional float
	}

	if dbSite.Address.Valid {
		siteInfo.Address = dbSite.Address.String
	} else {
		siteInfo.Address = "" // Default for optional string
	}

	return siteInfo, nil
}

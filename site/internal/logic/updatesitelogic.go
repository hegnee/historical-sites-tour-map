package logic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"site/internal/model"
	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// 1. Input Validation
	if in.SiteId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "SiteID cannot be empty")
	}

	// 2. Fetch Existing Record
	dbSite, err := l.svcCtx.SiteModel.FindOne(l.ctx, in.SiteId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "Site not found with ID: %s", in.SiteId)
		}
		l.Errorf("Database error fetching site for update. ID: %s, Error: %v", in.SiteId, err)
		return nil, status.Errorf(codes.Internal, "Database error fetching site: %v", err)
	}

	// 3. Update Fields selectively based on request
	updated := false

	if in.Name != "" && dbSite.Name != in.Name {
		dbSite.Name = in.Name
		updated = true
	}

	// Handle optional string fields: update if new value provided, or if clearing an existing one.
	if in.Description != dbSite.Description.String || (in.Description == "" && dbSite.Description.Valid) || (in.Description != "" && !dbSite.Description.Valid) {
		dbSite.Description = sql.NullString{String: in.Description, Valid: in.Description != ""}
		updated = true
	}
	if in.Category != dbSite.Category.String || (in.Category == "" && dbSite.Category.Valid) || (in.Category != "" && !dbSite.Category.Valid) {
		dbSite.Category = sql.NullString{String: in.Category, Valid: in.Category != ""}
		updated = true
	}
	if in.Region != dbSite.Region.String || (in.Region == "" && dbSite.Region.Valid) || (in.Region != "" && !dbSite.Region.Valid) {
		dbSite.Region = sql.NullString{String: in.Region, Valid: in.Region != ""}
		updated = true
	}
	if in.HistoricalPeriod != dbSite.HistoricalPeriod.String || (in.HistoricalPeriod == "" && dbSite.HistoricalPeriod.Valid) || (in.HistoricalPeriod != "" && !dbSite.HistoricalPeriod.Valid) {
		dbSite.HistoricalPeriod = sql.NullString{String: in.HistoricalPeriod, Valid: in.HistoricalPeriod != ""}
		updated = true
	}
	if in.Address != dbSite.Address.String || (in.Address == "" && dbSite.Address.Valid) || (in.Address != "" && !dbSite.Address.Valid) {
		dbSite.Address = sql.NullString{String: in.Address, Valid: in.Address != ""}
		updated = true
	}

	// Handle optional float fields
	if dbSite.Latitude.Float64 != in.Latitude || !dbSite.Latitude.Valid {
		dbSite.Latitude = sql.NullFloat64{Float64: in.Latitude, Valid: true} // Assuming if lat/lon is in req, it's intended to be set
		updated = true
	}
	if dbSite.Longitude.Float64 != in.Longitude || !dbSite.Longitude.Valid {
		dbSite.Longitude = sql.NullFloat64{Float64: in.Longitude, Valid: true} // Assuming if lat/lon is in req, it's intended to be set
		updated = true
	}

	if updated {
		dbSite.UpdatedAt = time.Now()
		// 4. Update Database (if changes were made)
		err = l.svcCtx.SiteModel.Update(l.ctx, dbSite)
		if err != nil {
			l.Errorf("Failed to update site in database. ID: %s, Error: %v", in.SiteId, err)
			return nil, status.Errorf(codes.Internal, "Failed to update site: %v", err)
		}
	}

	// 5. Construct Response from potentially updated dbSite
	siteInfo := &pb.SiteInfo{
		Id:        dbSite.Id,
		Name:      dbSite.Name,
		CreatedAt: dbSite.CreatedAt.Unix(),
		UpdatedAt: dbSite.UpdatedAt.Unix(), // This will be the new time if updated, or old time if not
	}

	if dbSite.Description.Valid {
		siteInfo.Description = dbSite.Description.String
	}
	if dbSite.Category.Valid {
		siteInfo.Category = dbSite.Category.String
	}
	if dbSite.Region.Valid {
		siteInfo.Region = dbSite.Region.String
	}
	if dbSite.HistoricalPeriod.Valid {
		siteInfo.HistoricalPeriod = dbSite.HistoricalPeriod.String
	}
	if dbSite.Latitude.Valid {
		siteInfo.Latitude = dbSite.Latitude.Float64
	}
	if dbSite.Longitude.Valid {
		siteInfo.Longitude = dbSite.Longitude.Float64
	}
	if dbSite.Address.Valid {
		siteInfo.Address = dbSite.Address.String
	}

	// 6. Return Response
	return siteInfo, nil
}

package logic

import (
	"context"
	"database/sql"
	"time"

	"site/internal/model"
	"site/internal/svc"
	"site/pb"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// 1. Input Validation
	if in.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Site name cannot be empty")
	}

	// 2. Prepare Data for Insertion
	newUUID, err := uuid.NewRandom()
	if err != nil {
		l.Errorf("Failed to generate UUID for new site: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to generate site ID: %v", err)
	}
	siteID := newUUID.String()
	now := time.Now()

	newSiteData := model.Sites{
		Id:   siteID,
		Name: in.Name,
		Description: sql.NullString{
			String: in.Description,
			Valid:  in.Description != "",
		},
		Category: sql.NullString{
			String: in.Category,
			Valid:  in.Category != "",
		},
		Region: sql.NullString{
			String: in.Region,
			Valid:  in.Region != "",
		},
		HistoricalPeriod: sql.NullString{
			String: in.HistoricalPeriod,
			Valid:  in.HistoricalPeriod != "",
		},
		Latitude: sql.NullFloat64{
			Float64: in.Latitude,
			Valid:   true, // Assuming latitude is always provided if part of req.
		},
		Longitude: sql.NullFloat64{
			Float64: in.Longitude,
			Valid:   true, // Assuming longitude is always provided if part of req.
		},
		Address: sql.NullString{
			String: in.Address,
			Valid:  in.Address != "",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	// 3. Insert into Database
	// The Insert method might return sql.Result or an error.
	// We typically don't need the sql.Result for this operation if ID is client-generated or timestamps are set.
	_, err = l.svcCtx.SiteModel.Insert(l.ctx, &newSiteData)
	if err != nil {
		l.Errorf("Failed to create site in database. Name: %s, Error: %v", in.Name, err)
		return nil, status.Errorf(codes.Internal, "Failed to create site: %v", err)
	}

	// 4. Construct Response
	siteInfo := &pb.SiteInfo{
		Id:               newSiteData.Id,
		Name:             newSiteData.Name,
		Description:      newSiteData.Description.String, // Will be empty if not Valid
		Category:         newSiteData.Category.String,    // Will be empty if not Valid
		Region:           newSiteData.Region.String,      // Will be empty if not Valid
		HistoricalPeriod: newSiteData.HistoricalPeriod.String, // Will be empty if not Valid
		Latitude:         newSiteData.Latitude.Float64,   // Will be 0.0 if not Valid (though we set Valid:true)
		Longitude:        newSiteData.Longitude.Float64,  // Will be 0.0 if not Valid (though we set Valid:true)
		Address:          newSiteData.Address.String,     // Will be empty if not Valid
		CreatedAt:        newSiteData.CreatedAt.Unix(),
		UpdatedAt:        newSiteData.UpdatedAt.Unix(),
	}
	// For fields where Valid was explicitly false (e.g. if Description was empty),
	// .String will return "" which is fine for proto.
	// For floats where Valid was true, .Float64 returns the value.

	// 5. Return Response
	return siteInfo, nil
}

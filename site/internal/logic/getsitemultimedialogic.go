package logic

import (
	"context"
	// "database/sql" // Import if model.Multimedia fields are sql.NullString etc. and used directly

	"site/internal/model" // For model.Multimedia type
	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// 1. Input Validation
	if in.SiteId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "SiteID cannot be empty")
	}

	// 2. Query Database using sqlx
	conn := l.svcCtx.MultimediaModel.Conn() // Assuming MultimediaModel is available in svcCtx
	multimediaTableName := "multimedia"     // Hardcoded table name

	query := "SELECT id, site_id, type, url, description, created_at, updated_at FROM " + multimediaTableName + " WHERE site_id = ? ORDER BY created_at ASC"
	args := []interface{}{in.SiteId}

	// 3. Execute Data Query
	multimediaDB := []*model.Multimedia{} // Slice of pointers to the model struct
	err := conn.QueryRowsCtx(l.ctx, &multimediaDB, query, args...)
	if err != nil {
		l.Errorf("Failed to execute data query for site multimedia. SiteID: %s, Query: %s, Args: %v, Error: %v", in.SiteId, query, args, err)
		return nil, status.Errorf(codes.Internal, "Database error fetching multimedia: %v", err)
	}

	// 4. Map Results to Protobuf
	pbMultimediaList := make([]*pb.MultimediaInfo, 0, len(multimediaDB))
	for _, dbEntry := range multimediaDB {
		multimediaInfo := &pb.MultimediaInfo{
			Id:        dbEntry.Id,
			SiteId:    dbEntry.SiteId,
			Type:      dbEntry.Type,
			Url:       dbEntry.Url,
			CreatedAt: dbEntry.CreatedAt.Unix(),
			UpdatedAt: dbEntry.UpdatedAt.Unix(),
		}
		// Handle Description (assuming it's sql.NullString in model.Multimedia)
		if dbEntry.Description.Valid {
			multimediaInfo.Description = dbEntry.Description.String
		} else {
			multimediaInfo.Description = "" // Default for optional string
		}
		pbMultimediaList = append(pbMultimediaList, multimediaInfo)
	}

	// 5. Construct and Return Response
	return &pb.GetSiteMultimediaResp{Multimedia: pbMultimediaList}, nil
}

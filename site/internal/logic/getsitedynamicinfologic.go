package logic

import (
	"context"
	"database/sql" // For sql.ErrNoRows and sql.NullString
	"encoding/json"
	"errors"

	"site/internal/model" // For model.DynamicInfo type and model.ErrNotFound
	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// 1. Input Validation
	if in.SiteId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "SiteID cannot be empty")
	}

	// 2. Fetch Record from Database
	var dbEntry model.DynamicInfo // Assuming model.DynamicInfo is the struct type
	var err error

	// Attempt to use a specific method if available (e.g., l.svcCtx.DynamicInfoModel.FindOneBySiteId)
	// This is a placeholder for the check; actual detection of method existence in Go is complex
	// and usually relies on interface design or type assertions.
	// For this example, we'll assume that if DynamicInfoModel is not nil, we try a conventional method name.
	// If not, we fall back to raw sqlx. This logic might need refinement based on actual model capabilities.

	// A common pattern in go-zero is for FindOneBy<FieldName> methods to be generated if the field is unique.
	// Let's assume such a method `FindOneBySiteId` exists on `l.svcCtx.DynamicInfoModel`.
	// If it doesn't, the code would need to be structured to use the sqlx fallback.
	// For this exercise, we'll simulate this by attempting the direct method.
	// If `l.svcCtx.DynamicInfoModel.FindOneBySiteId` were a real method:
	// tempDbEntry, findErr := l.svcCtx.DynamicInfoModel.FindOneBySiteId(l.ctx, in.SiteId)
	// if findErr == nil && tempDbEntry != nil {
	//    dbEntry = *tempDbEntry // Dereference if it returns a pointer
	//    err = nil
	// } else {
	//    err = findErr // Keep the error to check for ErrNotFound etc.
	// }
	// Since we cannot dynamically check for method existence easily and the prompt suggests a fallback:
	// We will prioritize using the sqlx approach directly as it's more explicit for this context.

	conn := l.svcCtx.DynamicInfoModel.Conn()
	dynamicInfoTableName := "dynamic_info"
	query := "SELECT id, site_id, opening_hours, ticket_price, notices, updated_at FROM " + dynamicInfoTableName + " WHERE site_id = ?"
	args := []interface{}{in.SiteId}

	// QueryRowCtx directly into the fields of dbEntry
	// Need to scan into the individual fields because dbEntry is not a slice for QueryRowCtx
	// Let's adjust to scan into a temporary struct or directly if model fields are exported and match.
	// The `sqlx.Sqlconn.QueryRowCtx` expects a struct with `db` tags.
	// `model.DynamicInfo` should have these.
	err = conn.QueryRowCtx(l.ctx, &dbEntry, query, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, model.ErrNotFound) { // model.ErrNotFound might wrap sql.ErrNoRows
			return nil, status.Errorf(codes.NotFound, "Dynamic info not found for SiteID: %s", in.SiteId)
		}
		l.Errorf("Database error fetching dynamic info for SiteID %s: %v. Query: %s", in.SiteId, err, query)
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	// 3. Map Results to Protobuf
	pbDynamicInfo := &pb.DynamicInfo{
		SiteId: dbEntry.SiteId, // Assuming SiteId in dbEntry is non-nullable string
	}

	// Assuming UpdatedAt in dbEntry is non-nullable time.Time
	pbDynamicInfo.UpdatedAt = dbEntry.UpdatedAt.Unix()

	// Handle OpeningHours (assuming sql.NullString)
	if dbEntry.OpeningHours.Valid {
		pbDynamicInfo.OpeningHours = dbEntry.OpeningHours.String
	} else {
		pbDynamicInfo.OpeningHours = "" // Default for optional string
	}

	// Handle TicketPrice (assuming sql.NullString)
	if dbEntry.TicketPrice.Valid {
		pbDynamicInfo.TicketPrice = dbEntry.TicketPrice.String
	} else {
		pbDynamicInfo.TicketPrice = "" // Default for optional string
	}

	// Handle Notices (assuming sql.NullString containing JSON array)
	pbDynamicInfo.Notices = []string{} // Initialize to empty slice
	if dbEntry.Notices.Valid && dbEntry.Notices.String != "" {
		jsonErr := json.Unmarshal([]byte(dbEntry.Notices.String), &pbDynamicInfo.Notices)
		if jsonErr != nil {
			l.Warnf("Failed to unmarshal notices JSON for SiteID %s: %v. JSON: '%s'", in.SiteId, jsonErr, dbEntry.Notices.String)
			// Decide on behavior: return error, or return empty notices / partial data
			// For now, returning empty notices as per prompt's suggestion (pbDynamicInfo.Notices = []string{})
		}
	}

	// 4. Return Response
	return pbDynamicInfo, nil
}

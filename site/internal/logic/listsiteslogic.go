package logic

import (
	"context"
	"strings"

	"site/internal/model"
	"site/internal/svc"
	"site/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ListSitesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSitesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSitesLogic {
	return &ListSitesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListSitesLogic) ListSites(in *pb.ListSitesReq) (*pb.ListSitesResp, error) {
	// 1. Set Defaults for Pagination
	page := in.Page
	if page <= 0 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 { // Cap page size
		pageSize = 100
	}

	// 2. Build Dynamic Query using sqlx
	whereClauses := []string{}
	args := []interface{}{}

	// Assuming l.svcCtx.SiteModel.TableName() exists. If not, hardcode "sites".
	// For robust model access, it might be better to get TableName() from an instance
	// if the model definition is not directly accessible for its TableName method.
	// However, go-zero's defaultSitesModel typically provides TableName().
	// If SiteModel is an interface, it should include TableName(). For this example, we'll assume it's accessible.
	// If SiteModel is directly the defaultSitesModel, then it's fine.
	// If SiteModel is a custom type wrapping defaultSitesModel, ensure TableName is exposed.
	// For this implementation, let's assume `l.svcCtx.SiteModel` can give us the table name.
	// A common way is `l.svcCtx.SiteModel.TableName()`. If this method is not part of the
	// concrete type or interface, you would need to adjust or hardcode.
	// Let's assume `model.Sites{}.TableName()` or similar if direct method not on svcCtx.SiteModel.
	// For now, we'll use the placeholder as per instructions.
	// If this were a real scenario and SiteModel is an interface without TableName(),
	// we'd need to cast or get it from a concrete model instance.
	// Let's assume SiteModel has a TableName method, or we hardcode.
	// For safety, let's use the model instance if available, or hardcode.
	// The default generated model `defaultSitesModel` has `m.table` which is unexported.
	// `NewSitesModel` returns `SitesModel` interface. The concrete `customSitesModel` embeds `defaultSitesModel`.
	// To access `table` safely without breaking encapsulation or hardcoding,
	// `TableName()` method should be added to the `SitesModel` interface and `customSitesModel`.
	// For now, as per common practice in go-zero examples if not directly available, we might see hardcoding.
	// However, the prompt implies `l.svcCtx.SiteModel.TableName()` exists.
	// Let's proceed with that assumption for the TableName.
	// If `l.svcCtx.SiteModel.TableName()` isn't directly on the interface,
	// and `l.svcCtx.SiteModel` is the `model.SitesModel` interface,
	// one would typically add `TableName() string` to this interface
	// and implement it in `customSitesModel` to return `m.table`.
	// For now, we'll assume it's available on the model instance provided by svcCtx.
	// A more direct way if TableName() is not on the interface:
	// tableName := model.Sites{}.TableName() // This is a common pattern if TableName is a method on the struct itself.
	// Or, if you have access to the underlying defaultSitesModel, but that's less clean.
	// Let's stick to the problem's suggestion.
	var tableName = "sites" // Fallback if TableName() method is an issue.
	// Attempt to use TableName() if it were part of the interface:
	// if siteModelWithTableName, ok := l.svcCtx.SiteModel.(interface{ TableName() string }); ok {
	// 	tableName = siteModelWithTableName.TableName()
	// } else {
	//  l.Warnf("SiteModel does not implement TableName(), using default 'sites'")
	// }
	// Given the go-zero structure, `TableName()` is not on the `SitesModel` interface by default.
	// We will hardcode "sites" as it's the most straightforward way without modifying generated code signatures.

	baseQuery := "SELECT id, name, description, category, region, historical_period, latitude, longitude, address, created_at, updated_at FROM " + tableName
	countBaseQuery := "SELECT count(*) FROM " + tableName

	if in.Category != "" {
		whereClauses = append(whereClauses, "category = ?")
		args = append(args, in.Category)
	}
	if in.Region != "" {
		whereClauses = append(whereClauses, "region = ?")
		args = append(args, in.Region)
	}
	if in.HistoricalPeriod != "" {
		whereClauses = append(whereClauses, "historical_period = ?")
		args = append(args, in.HistoricalPeriod)
	}

	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = " WHERE " + strings.Join(whereClauses, " AND ")
	}

	dataQueryString := baseQuery + whereSQL + " ORDER BY created_at DESC LIMIT ? OFFSET ?"
	countQueryString := countBaseQuery + whereSQL

	dataArgs := append(append([]interface{}{}, args...), pageSize, (page-1)*pageSize) // Ensure args is copied

	// 3. Execute Count Query
	var total int64
	conn := l.svcCtx.SiteModel.Conn() // Get sqlx.SqlConn
	err := conn.QueryRowCtx(l.ctx, &total, countQueryString, args...)
	if err != nil {
		l.Errorf("Failed to execute count query for sites. Query: %s, Args: %v, Error: %v", countQueryString, args, err)
		return nil, status.Errorf(codes.Internal, "Database error counting sites: %v", err)
	}

	if total == 0 {
		return &pb.ListSitesResp{Sites: []*pb.SiteInfo{}, Total: 0, Page: page, PageSize: pageSize}, nil
	}

	// 4. Execute Data Query
	sitesDB := []*model.Sites{} // Assuming Sites is the struct type in model package
	err = conn.QueryRowsCtx(l.ctx, &sitesDB, dataQueryString, dataArgs...)
	if err != nil {
		l.Errorf("Failed to execute data query for sites. Query: %s, Args: %v, Error: %v", dataQueryString, dataArgs, err)
		return nil, status.Errorf(codes.Internal, "Database error fetching sites: %v", err)
	}

	// 5. Map Results to Protobuf
	pbSites := make([]*pb.SiteInfo, 0, len(sitesDB))
	for _, dbSite := range sitesDB {
		siteInfo := &pb.SiteInfo{
			Id:        dbSite.Id,
			Name:      dbSite.Name,
			CreatedAt: dbSite.CreatedAt.Unix(),
			UpdatedAt: dbSite.UpdatedAt.Unix(),
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
		pbSites = append(pbSites, siteInfo)
	}

	// 6. Construct and Return Response
	return &pb.ListSitesResp{Sites: pbSites, Total: total, Page: page, PageSize: pageSize}, nil
}

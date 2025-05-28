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

type SearchSitesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSitesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSitesLogic {
	return &SearchSitesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSitesLogic) SearchSites(in *pb.SearchSitesReq) (*pb.SearchSitesResp, error) {
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

	// 2. Input Validation
	trimmedKeyword := strings.TrimSpace(in.Keyword)
	if trimmedKeyword == "" {
		return &pb.SearchSitesResp{Sites: []*pb.SiteInfo{}, Total: 0, Page: page, PageSize: pageSize}, nil
	}

	// 3. Build Dynamic Query using sqlx (with FULLTEXT search)
	conn := l.svcCtx.SiteModel.Conn()
	tableName := "sites" // Hardcoded for simplicity

	baseQuery := "SELECT id, name, description, category, region, historical_period, latitude, longitude, address, created_at, updated_at FROM " + tableName
	countBaseQuery := "SELECT count(*) FROM " + tableName

	// Construct the WHERE clause for FULLTEXT search
	whereSQL := " WHERE MATCH(name, description) AGAINST(? IN NATURAL LANGUAGE MODE)"
	args := []interface{}{trimmedKeyword} // Use trimmed keyword

	// Final queries
	// For relevance ordering, you might use:
	// dataQueryString := baseQuery + whereSQL + " ORDER BY MATCH(name, description) AGAINST(? IN NATURAL LANGUAGE MODE) DESC LIMIT ? OFFSET ?"
	// And then dataArgs would need the keyword repeated: append([]interface{}{trimmedKeyword}, args..., pageSize, (page-1)*pageSize)
	// For simplicity, sticking to created_at DESC as per instructions.
	dataQueryString := baseQuery + whereSQL + " ORDER BY created_at DESC LIMIT ? OFFSET ?"
	countQueryString := countBaseQuery + whereSQL

	// Prepare arguments for data query
	dataArgs := append(append([]interface{}{}, args...), pageSize, (page-1)*pageSize) // Ensure args is copied

	// 4. Execute Count Query
	var total int64
	err := conn.QueryRowCtx(l.ctx, &total, countQueryString, args...)
	if err != nil {
		l.Errorf("Failed to execute count query for site search. Keyword: '%s', Query: %s, Args: %v, Error: %v", trimmedKeyword, countQueryString, args, err)
		return nil, status.Errorf(codes.Internal, "Database error counting sites for search: %v", err)
	}

	if total == 0 {
		return &pb.SearchSitesResp{Sites: []*pb.SiteInfo{}, Total: 0, Page: page, PageSize: pageSize}, nil
	}

	// 5. Execute Data Query (if total > 0)
	sitesDB := []*model.Sites{}
	err = conn.QueryRowsCtx(l.ctx, &sitesDB, dataQueryString, dataArgs...)
	if err != nil {
		l.Errorf("Failed to execute data query for site search. Keyword: '%s', Query: %s, Args: %v, Error: %v", trimmedKeyword, dataQueryString, dataArgs, err)
		return nil, status.Errorf(codes.Internal, "Database error fetching sites for search: %v", err)
	}

	// 6. Map Results to Protobuf
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

	// 7. Construct and Return Response
	return &pb.SearchSitesResp{Sites: pbSites, Total: total, Page: page, PageSize: pageSize}, nil
}

package svc

import (
	"site/internal/config"
	"site/internal/model" // Import the model package

	"github.com/zeromicro/go-zero/core/stores/sqlx" // Import for sqlx.SqlConn
)

type ServiceContext struct {
	Config           config.Config
	SiteModel        model.SitesModel        // Existing (assuming it will be added if not present)
	MultimediaModel  model.MultimediaModel   // New
	DynamicInfoModel model.DynamicInfoModel  // New
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Initialize sqlx.SqlConn (common for all models)
	// This assumes c.DB.DataSource is the DSN string
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:    c,
		SiteModel: model.NewSitesModel(sqlConn, c.Cache), // Assuming NewSitesModel and c.Cache are correct
		MultimediaModel:  model.NewMultimediaModel(sqlConn, c.Cache), // New
		DynamicInfoModel: model.NewDynamicInfoModel(sqlConn, c.Cache), // New
	}
}

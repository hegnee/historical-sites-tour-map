package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SitesModel = (*customSitesModel)(nil)

type (
	// SitesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSitesModel.
	SitesModel interface {
		sitesModel
	}

	customSitesModel struct {
		*defaultSitesModel
	}
)

// NewSitesModel returns a model for the database table.
func NewSitesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SitesModel {
	return &customSitesModel{
		defaultSitesModel: newSitesModel(conn, c, opts...),
	}
}

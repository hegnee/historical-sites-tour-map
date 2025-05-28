package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DynamicInfoModel = (*customDynamicInfoModel)(nil)

type (
	// DynamicInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDynamicInfoModel.
	DynamicInfoModel interface {
		dynamicInfoModel
	}

	customDynamicInfoModel struct {
		*defaultDynamicInfoModel
	}
)

// NewDynamicInfoModel returns a model for the database table.
func NewDynamicInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DynamicInfoModel {
	return &customDynamicInfoModel{
		defaultDynamicInfoModel: newDynamicInfoModel(conn, c, opts...),
	}
}

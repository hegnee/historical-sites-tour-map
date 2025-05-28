package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MultimediaModel = (*customMultimediaModel)(nil)

type (
	// MultimediaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMultimediaModel.
	MultimediaModel interface {
		multimediaModel
	}

	customMultimediaModel struct {
		*defaultMultimediaModel
	}
)

// NewMultimediaModel returns a model for the database table.
func NewMultimediaModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MultimediaModel {
	return &customMultimediaModel{
		defaultMultimediaModel: newMultimediaModel(conn, c, opts...),
	}
}

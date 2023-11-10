package repository

import (
	"gitlab.et-ns.net/connect/graph-ql-api/internal/adapter"
	"gorm.io/gorm"
)

func GetAdaptor() *gorm.DB {
	db := adapter.GetDb()

	return db
}

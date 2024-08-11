package model

import (
	"gorm.io/gorm"
)

type CommonDeleteResponse struct {
	Deleted    bool   `json:"deleted"`
	EntityName string `json:"entity_name"`
	IDDeleted  int    `json:"id_deleted"`
}

func CreateCommonDeleteResponse(EntityName string, IDDeleted int) *CommonDeleteResponse {
	return &CommonDeleteResponse{
		Deleted:    true,
		EntityName: EntityName,
		IDDeleted:  IDDeleted,
	}
}

func Paginate(pageSize, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

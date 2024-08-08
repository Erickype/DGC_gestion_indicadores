package model

import "gorm.io/gorm/schema"

type TableWithSchemaName interface {
	TableName(name schema.Namer) string
}

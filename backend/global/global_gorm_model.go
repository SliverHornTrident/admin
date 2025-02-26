//go:build (gorm || gorms) && (tidb || mysql || postgres || sqlite || clickhouse || mssql || sqlserver || oracle)

package global

import (
	soft "gorm.io/plugin/soft_delete"
	"time"
)

// Model 默认模型
type Model struct {
	ID        uint            `json:"ID" gorm:"primaryKey;column:id;comment:主键ID"`
	CreatedAt time.Time       `json:"CreatedAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time       `json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt time.Time       `json:"-" gorm:"default:null;column:deleted_at;comment:删除时间"`
	IsDelete  *soft.DeletedAt `json:"-" gorm:"index;softDelete:flag,DeletedAtField:DeletedAt;default:0;comment:删除标志"`
}

//go:build (gorm || gorms) && (tidb || mysql || postgres || sqlite || clickhouse || mssql || sqlserver || oracle) && sonyflake

package global

import (
	"github.com/SliverHornTrident/components/component"
	"gorm.io/gorm"
	soft "gorm.io/plugin/soft_delete"
	"time"
)

// ModelSonyflake 默认雪花模型
type ModelSonyflake struct {
	ID        uint64          `json:"ID" gorm:"primaryKey;autoIncrement:false;column:id;comment:主键ID"`
	CreatedAt time.Time       `json:"CreatedAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time       `json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt time.Time       `json:"-" gorm:"default:null;column:deleted_at;comment:删除时间"`
	IsDelete  *soft.DeletedAt `json:"-" gorm:"index;softDelete:flag,DeletedAtField:DeletedAt;default:0;comment:删除标志"`
}

func (m *ModelSonyflake) BeforeCreate(tx *gorm.DB) error {
	id, err := component.Sonyflake.NextID()
	if err != nil {
		return err
	}
	m.ID = id
	return nil
}

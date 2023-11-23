package base

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Model struct {
	ID        uint64                `gorm:"AUTO_INCREMENT;primaryKey;comment:'主键'" json:"id"`
	CreatedAt *time.Time            `gorm:"column:created_at;comment:'创建时间'" json:"created_at"`
	UpdatedAt *time.Time            `gorm:"column:updated_at;comment:'更新时间'" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;uniqueIndex:unique_index;comment:非空表示已删除;default:0" json:"-"`
}

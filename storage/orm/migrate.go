package orm

import "time"

// Migration 数据迁移
type Migration struct {
	Version   string    `gorm:"primaryKey"`
	ApplyTime time.Time `gorm:"autoCreateTime"`
}

// TableName 表名
func (Migration) TableName() string {
	return "migration"
}

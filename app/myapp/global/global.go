package global

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// Log is the global logger
	Log *zap.Logger
	// CfgMap is the global config file
	CfgMap = map[string]string{
		"development": "/app/myapp/etc/config.yml",
		"production":  "/app/myapp/etc/config_production.yml",
		"test":        "/app/myapp/etc/config.yml",
		"":            "/app/myapp/etc/config_production.yml",
	}
)

// COMMODEL is the common model for all models.
type COMMODEL struct {
	ID        int64          `gorm:"primarykey"`                             // 主键ID
	CreatedAt time.Time      `gorm:"index:idx_created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"index:idx_updated_at" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at" json:"-"`          // 删除时间
}

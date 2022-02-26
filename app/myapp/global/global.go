package global

import (
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
	}

	// DB is the global database
	DB  *gorm.DB
	err error
)

package database

import (
	"log"
	"time"

	"github.com/elton/project-layout/app/myapp/global"
	"github.com/elton/project-layout/config"
	"github.com/elton/project-layout/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB is the database connection
	DB *gorm.DB
)

func init() {
	DB, _ = InitDatabase()
}

// InitDatabase initial the database
func InitDatabase() (*gorm.DB, error) {
	// Read configuration file.
	if err := config.ReadConfig(global.CfgMap); err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(mysql.Open(config.AppCfg.Database.Dsn), &gorm.Config{})
	if err != nil {
		logger.Sugar.Errorf("Unable to connect to database: %s", err.Error())
		return nil, err
	}
	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.Sugar.Infof("Connected to database: %s", config.AppCfg.Database.Dsn)
	return db, nil
}

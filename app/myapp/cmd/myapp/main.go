package main

import (
	"github.com/elton/project-layout/app/myapp/internal/app/models"
	"github.com/elton/project-layout/app/myapp/internal/pkg/database"
	"github.com/elton/project-layout/pkg/logger"
)

func main() {
	if err := database.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}); err != nil {
		logger.Sugar.Errorf("Migrate failed: %v", err)
	}
	logger.Sugar.Debugf("Migrate success")

	server := InitServer()
	server.Start()
}

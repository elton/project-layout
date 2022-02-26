package main

import (
	"github.com/elton/project-layout/app/myapp/internal/model"
	"github.com/elton/project-layout/app/myapp/internal/pkg/database"
	"github.com/elton/project-layout/app/myapp/internal/pkg/server"
	"github.com/elton/project-layout/app/myapp/internal/service"
	"github.com/elton/project-layout/pkg/logger"
)

func main() {
	greeting, err := service.Hello()
	if err != nil {
		panic(err)
	}
	logger.Sugar.Debugf("%s", greeting)
	if err := database.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}); err != nil {
		logger.Sugar.Errorf("Migrate failed: %v", err)
	}
	logger.Sugar.Debugf("Migrate success")

	server.Start()

}

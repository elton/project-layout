package main

import (
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
	server.Start()
	database.InitDatabase()
}

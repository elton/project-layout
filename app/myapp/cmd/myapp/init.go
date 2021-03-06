package main

import (
	"github.com/elton/project-layout/app/myapp/api/controllers"
	"github.com/elton/project-layout/app/myapp/api/router"
	"github.com/elton/project-layout/app/myapp/internal/app/models"
	"github.com/elton/project-layout/app/myapp/internal/app/service"
	"github.com/elton/project-layout/app/myapp/internal/pkg/database"
	"github.com/elton/project-layout/app/myapp/internal/pkg/server"
)

// InitServer initializes the server
func InitServer() *server.Server {
	app := server.NewApp()
	db := database.DB
	iUserRepo := models.NewUserRepository(db)
	iUserService := service.NewUserService(iUserRepo)
	userController := controllers.NewUserController(iUserService)
	router := router.NewRouter(userController)
	server := server.NewServer(app, router)
	return server
}

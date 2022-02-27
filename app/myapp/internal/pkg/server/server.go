package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/elton/project-layout/app/myapp/api/router"
	"github.com/elton/project-layout/config"
	"github.com/gofiber/fiber/v2"
)

// Server represents the webserver
type Server struct {
	app    *fiber.App
	router *router.Router
}

// NewServer creates a new webserver
func NewServer(app *fiber.App, router *router.Router) *Server {
	return &Server{
		app:    app,
		router: router,
	}
}

// Start setup a webserver
func (s *Server) Start() {

	s.router.With(s.app)

	go func() {
		log.Fatal(s.app.Listen(config.AppCfg.Server.Port))
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = s.app.Shutdown()
}

package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/elton/project-layout/app/myapp/internal/router"
	"github.com/elton/project-layout/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const configPath = "/app/myapp/etc/config.yml"

// Start setup a webserver
func Start() {
	// perfork, _ := strconv.ParseBool(config.Config("PREFORK"))
	if err := config.ReadConfig(configPath); err != nil {
		log.Fatal(err)
	}
	// Web server
	app := fiber.New(fiber.Config{
		Prefork:       config.AppCfg.Server.Prefork,
		StrictRouting: true,
		ServerHeader:  config.AppCfg.Server.Name,
		ReadTimeout:   time.Duration(config.AppCfg.Server.ReadTimeout) * time.Second,
		WriteTimeout:  time.Duration(config.AppCfg.Server.WriteTimeout) * time.Second,
	})

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "${pid} | ${ua} | ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
	}))
	app.Use(cors.New())

	router.InitializeRouters(app)

	go func() {
		if err := app.Listen(config.AppCfg.Server.Port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()
}

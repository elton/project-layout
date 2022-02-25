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
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
		Format: "${locals:requestid} | ${ua} | ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
	}))
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(csrf.New())

	router.InitializeRouters(app)

	go func() {
		log.Fatal(app.Listen(config.AppCfg.Server.Port))
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()
}

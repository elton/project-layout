package server

import (
	"log"
	"time"

	"github.com/elton/project-layout/app/myapp/global"
	"github.com/elton/project-layout/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// NewApp creates a new App instance
func NewApp() (app *fiber.App) {
	// Read configuration file.
	if err := configs.ReadConfig(global.CfgMap); err != nil {
		log.Fatal(err)
	}
	// Web server
	app = fiber.New(fiber.Config{
		Prefork:       configs.AppCfg.Server.Prefork,
		StrictRouting: true,
		ServerHeader:  configs.AppCfg.Server.Name,
		ReadTimeout:   time.Duration(configs.AppCfg.Server.ReadTimeout) * time.Second,
		WriteTimeout:  time.Duration(configs.AppCfg.Server.WriteTimeout) * time.Second,
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

	return
}

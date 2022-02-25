package pkg

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/elton/project-layout/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Start setup a webserver
func Start(port string) {
	perfork, _ := strconv.ParseBool(config.Config("PREFORK"))
	// Web server
	app := fiber.New(fiber.Config{
		Prefork:       perfork,
		StrictRouting: true,
		ServerHeader:  "BTOE-Server",
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
	})

	// Middleware
	// app.Use(middleware.Logger())
	app.Use(cors.New())

	// router.InitializeRouters(app)

	go func() {
		if err := app.Listen(port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()
}

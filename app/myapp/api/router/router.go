package router

import (
	"github.com/elton/project-layout/app/myapp/api/controllers"
	"github.com/gofiber/fiber/v2"
)

// Router represents the router of the application
type Router struct {
	user controllers.UserController
}

// NewRouter creates a new router
func NewRouter(user controllers.UserController) *Router {
	return &Router{
		user: user,
	}
}

// With initializes the router
func (r *Router) With(app *fiber.App) {

	// healthcheck
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", r.user.GetUserByName)

	// test panic recover
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	panic("I'm an error")
	// })
}

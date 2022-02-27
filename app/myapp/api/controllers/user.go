package controllers

import (
	"net/http"

	"github.com/elton/project-layout/app/myapp/internal/app/service"
	"github.com/elton/project-layout/pkg/logger"
	"github.com/elton/project-layout/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// UserController is a controller for user
type UserController struct {
	UserService service.IUserService
}

// NewUserController creates a new user controller
func NewUserController(userService service.IUserService) UserController {
	return UserController{
		UserService: userService,
	}
}

// GetUserByName returns a user by name
func (u *UserController) GetUserByName(c *fiber.Ctx) error {
	name := c.Query("name")
	logger.Sugar.Debugf("Get user by name: %s", name)
	user, err := u.UserService.GetUserByName(c.Context(), name)
	if err != nil {
		return err
	}
	response.ResultJSON(c, http.StatusOK, user)
	return nil
}

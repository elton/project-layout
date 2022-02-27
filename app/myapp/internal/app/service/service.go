package service

import (
	"context"

	"github.com/elton/project-layout/app/myapp/internal/app/models"
)

// IUserService presents a service for users.
type IUserService interface {
	GetUserByName(ctx context.Context, name string) (*models.User, error)
}

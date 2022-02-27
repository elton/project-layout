package service

import (
	"context"

	"github.com/elton/project-layout/app/myapp/internal/models"
)

type userService struct {
	userRepo models.IUserRepo
}

// NewUserService get a new user service.
func NewUserService(userRepo models.IUserRepo) IUserService {
	return &userService{userRepo}
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*models.User, error) {
	return u.userRepo.GetUserByName(ctx, name)
}

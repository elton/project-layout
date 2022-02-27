package models

import (
	"context"

	"gorm.io/gorm"
)

// IUserRepo return a user object by given query string
type IUserRepo interface {
	GetUserByName(ctx context.Context, name string) (*User, error)
}

type mySQLRepository struct {
	DB *gorm.DB
}

package models

import (
	"context"

	"gorm.io/gorm"
)

// IUserQuery return a user object by given query string
type IUserQuery interface {
	GetUserByName(ctx context.Context, name string) (*User, error)
}

type mySQLRepository struct {
	DB *gorm.DB
}

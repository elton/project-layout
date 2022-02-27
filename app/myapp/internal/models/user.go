package models

import (
	"context"

	"github.com/elton/project-layout/app/myapp/global"
	"github.com/elton/project-layout/pkg/logger"
	"gorm.io/gorm"
)

// User represents a user.
type User struct {
	global.COMMODEL
	Name   string `gorm:"index:idx_name" json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// NewUserRepository presents a repository for users.
func NewUserRepository(DB *gorm.DB) IUserQuery {
	return &mySQLRepository{DB}
}

func (m *mySQLRepository) GetUserByName(ctx context.Context, name string) (*User, error) {
	var user User
	if err := m.DB.Where("name = ?", name).First(&user).Error; err != nil {
		logger.Sugar.Errorf("failed to get user by name: %s", err)
		return nil, err
	}
	return &user, nil
}

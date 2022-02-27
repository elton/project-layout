package service

import (
	"context"
	"testing"

	"github.com/elton/project-layout/app/myapp/internal/mock"
	"github.com/elton/project-layout/app/myapp/internal/models"
	"github.com/golang/mock/gomock"
)

func TestGetUserByName(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := mock.NewMockIUserRepo(ctl)

	gomock.InOrder(
		mockRepo.EXPECT().GetUserByName(context.TODO(), "name").Return(&models.User{}, nil),
	)

	svc := NewUserService(mockRepo)

	user, err := svc.GetUserByName(context.TODO(), "name")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	t.Logf("user: %v", user)
}

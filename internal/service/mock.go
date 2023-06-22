package service

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/witalok2/test-dev-golang-worker/internal/entity"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateClient(ctx context.Context, client entity.Client) error {
	args := m.Called(ctx, client)
	return args.Error(0)
}

func (m *MockService) UpdateClient(ctx context.Context, client entity.Client) error {
	args := m.Called(ctx, client)
	return args.Error(0)
}

package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/witalok2/test-dev-golang-worker/internal/entity"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) CreateClient(ctx context.Context, client entity.Client) error {
	args := m.Called(ctx, client)
	return args.Error(0)
}

func (m *RepositoryMock) UpdateClient(ctx context.Context, client entity.Client) error {
	args := m.Called(ctx, client)
	return args.Error(0)
}

func (m *RepositoryMock) Close() {
	m.Called()
}

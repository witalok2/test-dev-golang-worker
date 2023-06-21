package service

import (
	"context"

	"github.com/witalok2/test-dev-golang-worker/internal/entity"
	repo "github.com/witalok2/test-dev-golang-worker/internal/repository"
)

type Service interface {
	CreateClient(ctx context.Context, client entity.Client) error
	UpdateClient(ctx context.Context, client entity.Client) error
}

type service struct {
	db repo.Repository
}

func NewService(ctx context.Context, db repo.Repository) *service {
	return &service{
		db: db,
	}
}

func (s service) CreateClient(ctx context.Context, client entity.Client) error {
	return s.db.CreateClient(ctx, client)
}

func (s service) UpdateClient(ctx context.Context, client entity.Client) error {
	return s.db.UpdateClient(ctx, client)
}

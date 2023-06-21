package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/test-dev-golang-worker/internal/entity"
)

type Repository interface {
	CreateClient(ctx context.Context, client entity.Client) error
	UpdateClient(ctx context.Context, client entity.Client) error
	Close()
}

type repository struct {
	db *sqlx.DB
}

func NewReaderConnection(URI string) (Repository, error) {
	db, err := sqlx.Connect("postgres", URI)
	if err != nil {
		return &repository{}, errors.New("failed to connect to read database")
	}

	return &repository{db}, nil
}

func (r *repository) CreateClient(ctx context.Context, client entity.Client) error {
	rows, err := r.db.QueryContext(ctx, SqlCreateClient,
		client.Name,
		client.LastName,
		client.Contact,
		client.Address,
		client.Birthday,
		client.CPF,
	)
	if err != nil {
		logger.WithError(err).Error("error creating client")
		return err
	}

	defer rows.Close()

	return nil
}

func (r *repository) UpdateClient(ctx context.Context, client entity.Client) error {
	_, err := r.db.ExecContext(ctx, SqlUpdateClient,
		client.ID,
		client.Name,
		client.LastName,
		client.Contact,
		client.Address,
		client.Birthday,
		client.CPF,
	)
	if err != nil {
		logger.WithError(err).Error("error update client")
		return err
	}

	return nil
}

func (r *repository) Close() {
	err := r.db.Close()
	if err != nil {
		logger.WithError(err).Error("error closing database connection")
	}
}

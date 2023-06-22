package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/witalok2/test-dev-golang-worker/internal/entity"
)

func TestCreateClient(t *testing.T) {
	repo := new(RepositoryMock)

	client := entity.Client{
		Name:     "John Doe",
		LastName: "Doe",
		Contact:  "john.doe@example.com",
		Address:  "123 Main St",
		Birthday: "1994-02-04",
		CPF:      "123456789",
	}

	repo.On("CreateClient", mock.Anything, client).Return(nil)

	ctx := context.Background()

	err := repo.CreateClient(ctx, client)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestUpdateClient(t *testing.T) {
	repo := new(RepositoryMock)

	client := entity.Client{
		ID:       uuid.New(),
		Name:     "John Doe",
		LastName: "Doe",
		Contact:  "john.doe@example.com",
		Address:  "123 Main St",
		Birthday: "1994-02-04",
		CPF:      "123456789",
	}

	repo.On("UpdateClient", mock.Anything, client).Return(nil)

	ctx := context.Background()

	err := repo.UpdateClient(ctx, client)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestClose(t *testing.T) {
	repo := new(RepositoryMock)

	repo.On("Close").Return()

	repo.Close()

	repo.AssertExpectations(t)
}

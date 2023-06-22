package service

import (
	"context"
	"sort"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/witalok2/test-dev-golang-worker/internal/entity"
	repo "github.com/witalok2/test-dev-golang-worker/internal/repository"
)

type ServiceTestSuite struct {
	suite.Suite
	ctx context.Context

	databaseMock *repo.RepositoryMock

	service *service
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.databaseMock = new(repo.RepositoryMock)
	suite.service = NewService(suite.ctx, suite.databaseMock)
}

func (suite *ServiceTestSuite) TestCreateClient() {
	cases := map[string]struct {
		client entity.Client
		err    error
	}{
		`create client`: {
			client: entity.Client{
				Name:     "Witalo",
				LastName: "Matheus",
			},
			err: nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.databaseMock.On("CreateClient", suite.ctx, cases[key].client).Return(nil)

			err := suite.service.CreateClient(suite.ctx, cases[key].client)

			suite.NoError(err)

			suite.databaseMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *ServiceTestSuite) TestUpdateClient() {
	cases := map[string]struct {
		client entity.Client
		err    error
	}{
		`update client`: {
			client: entity.Client{
				ID:       uuid.New(),
				Name:     "John Doe",
				LastName: "Smith",
			},
			err: nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.databaseMock.On("UpdateClient", suite.ctx, cases[key].client).Return(nil)

			err := suite.service.UpdateClient(suite.ctx, cases[key].client)

			suite.NoError(err)

			suite.databaseMock.AssertExpectations(suite.T())
		})
	}
}

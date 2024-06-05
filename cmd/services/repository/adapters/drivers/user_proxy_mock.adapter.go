package adapters

import (
	"HEXAGONAL-GO/cmd/services/repository"
	"HEXAGONAL-GO/cmd/services/repository/models"
	"context"
)

type UserProxyAdapterMock struct {
	ctx        context.Context
	repository *repository.Repository
}

func (a *UserProxyAdapterMock) GetUser(username string) (*models.User, error) {
	return &models.User{
		Username: "%username",
		Password: "%password",
		Deleted:  false,
	}, nil
}

func NewUserProxyAdapterMock(ctx context.Context, repository *repository.Repository) UserProxyAdapterMock {
	return UserProxyAdapterMock{
		ctx:        ctx,
		repository: repository,
	}
}

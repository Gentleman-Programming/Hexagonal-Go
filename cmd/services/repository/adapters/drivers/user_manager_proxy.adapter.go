package adapters

import (
	"HEXAGONAL-GO/cmd/services/repository"
	"HEXAGONAL-GO/cmd/services/repository/models"
	"context"
)

type UserManagerProxyAdapter struct {
	ctx        context.Context
	repository *repository.Repository
}

func (a *UserManagerProxyAdapter) GetUser(username string) (*models.User, error) {
	return a.repository.GetUser(a.ctx, username)
}

func NewUserManagerProxyAdapter(ctx context.Context, repository *repository.Repository) UserManagerProxyAdapter {
	return UserManagerProxyAdapter{
		ctx:        ctx,
		repository: repository,
	}
}

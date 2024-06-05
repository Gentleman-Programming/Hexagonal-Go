package adapters

import (
	"HEXAGONAL-GO/cmd/services/repository"
	"HEXAGONAL-GO/cmd/services/repository/models"
	"context"
)

type UserProxyAdapter struct {
	ctx        context.Context
	repository *repository.Repository
}

func (a *UserProxyAdapter) GetUser(username string) (*models.User, error) {
	return a.repository.GetUser(a.ctx, username)
}

func NewUserProxyAdapter(ctx context.Context, repository *repository.Repository) UserProxyAdapter {
	return UserProxyAdapter{
		ctx:        ctx,
		repository: repository,
	}
}

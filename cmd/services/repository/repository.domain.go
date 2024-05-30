package repository

import (
	"HEXAGONAL-GO/cmd/services/repository/models"
	"context"
)

type Repository struct {
	// driverns
}

func (r *Repository) GetUser(ctx context.Context, username string) (*models.User, error) {
	return &models.User{}, nil
}

func NewRepositoryService() *Repository {
	return &Repository{}
}

package repository

import (
	"HEXAGONAL-GO/cmd/services/repository/models"
)

type Repository struct {
	// driverns
}

func (r *Repository) GetUser(username string) (*models.User, error) {
	return &models.User{}, nil
}

func NewRepositoryService() *Repository {
	return &Repository{}
}

package ports

import (
	"HEXAGONAL-GO/cmd/services/repository/models"
)

type ForUser interface {
	GetUser(username string) (*models.User, error)
}

package ports

import (
	"HEXAGONAL-GO/cmd/services/dashboard-api/models"
)

type ForUser interface {
	GetUser(username string) (*models.User, error)
}

package ports

import (
	"HEXAGONAL-GO/cmd/services/repository/models"
)

type ForManagingUser interface {
	GetUser(username string) (*models.User, error)
}

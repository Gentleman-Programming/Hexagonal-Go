package ports

import "HEXAGONAL-GO/cmd/services/dashboard-api/models"

type ForQueryingUser interface {
	GetUser(username string) (*models.User, error)
}

package ports

import "HEXAGONAL-GO/cmd/services/dashboard-api/models"

type ForAuthenticatingUser interface {
	Authenticate(username string, password string) (*models.User, error)
}

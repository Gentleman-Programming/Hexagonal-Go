package ports

import "HEXAGONAL-GO/cmd/services/control-plane/models"

type ForManagingAuthentication interface {
	Authenticate(username string, password string) (*models.User, error)
}

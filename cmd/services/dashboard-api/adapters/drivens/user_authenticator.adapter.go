package adapters

import (
	control_plane_drivers "HEXAGONAL-GO/cmd/services/control-plane/ports/drivers"
	"HEXAGONAL-GO/cmd/services/dashboard-api/models"
	"context"
)

type UserAuthenticatorAdapter struct {
	ctx    context.Context
	driver control_plane_drivers.ForManagingAuthentication
}

func (a *UserAuthenticatorAdapter) Authenticate(username string, password string) (*models.User, error) {
	user, err := a.driver.Authenticate(username, password)
	if err != nil {
		return nil, err
	}

	mappedUser := models.User{
		Username: user.Username,
		Password: user.Password,
	}

	return &mappedUser, nil
}

func NewUserAuthenticatorAdapter(ctx context.Context, driver control_plane_drivers.ForManagingAuthentication) *UserAuthenticatorAdapter {
	return &UserAuthenticatorAdapter{
		ctx:    ctx,
		driver: driver,
	}
}

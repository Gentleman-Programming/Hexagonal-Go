package adapters

import (
	control_plane_drivers "HEXAGONAL-GO/cmd/services/control-plane/ports/drivers"
	"context"
)

type UserAuthorizerAdapter struct {
	ctx    context.Context
	driver control_plane_drivers.ForManagingAuthorization
}

func (a *UserAuthorizerAdapter) Authorize(userName string) (bool, error) {
	return a.driver.Authorize(userName)
}

func NewUserAuthorizerAdapter(ctx context.Context, driver control_plane_drivers.ForManagingAuthorization) *UserAuthorizerAdapter {
	return &UserAuthorizerAdapter{
		ctx:    ctx,
		driver: driver,
	}
}

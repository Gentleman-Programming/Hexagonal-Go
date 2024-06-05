package adapters

import (
	control_plane "HEXAGONAL-GO/cmd/services/control-plane"
	"HEXAGONAL-GO/cmd/services/control-plane/models"
	"context"
)

type AuthenticatorProxyAdapterMock struct {
	ctx          context.Context
	controlPlane *control_plane.ControlPlane
}

func (a AuthenticatorProxyAdapterMock) Authenticate(username string, password string) (*models.User, error) {
	return &models.User{
		Username: username,
		Password: password,
	}, nil
}

func NewAuthenticatorProxyAdapterMock(ctx context.Context, controlPlane *control_plane.ControlPlane) AuthenticatorProxyAdapterMock {
	return AuthenticatorProxyAdapterMock{
		ctx:          ctx,
		controlPlane: controlPlane,
	}
}

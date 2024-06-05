package adapters

import (
	control_plane "HEXAGONAL-GO/cmd/services/control-plane"
	"HEXAGONAL-GO/cmd/services/control-plane/models"
	"context"
)

type AuthenticatorProxyAdapter struct {
	ctx          context.Context
	controlPlane *control_plane.ControlPlane
}

func (a AuthenticatorProxyAdapter) Authenticate(username string, password string) (*models.User, error) {
	return a.controlPlane.Authenticate(username, password)
}

func NewAuthenticatorProxyAdapter(ctx context.Context, controlPlane *control_plane.ControlPlane) AuthenticatorProxyAdapter {
	return AuthenticatorProxyAdapter{
		ctx:          ctx,
		controlPlane: controlPlane,
	}
}

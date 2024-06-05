package adapters

import (
	control_plane "HEXAGONAL-GO/cmd/services/control-plane"
	"HEXAGONAL-GO/cmd/services/control-plane/models"
	"context"
)

type AuthenticationManagerProxyAdapter struct {
	ctx          context.Context
	controlPlane *control_plane.ControlPlane
}

func (a AuthenticationManagerProxyAdapter) Authenticate(username string, password string) (*models.User, error) {
	return a.controlPlane.Authenticate(a.ctx, username, password)
}

func NewAuthenticationManagerProxyAdapter(ctx context.Context, controlPlane *control_plane.ControlPlane) AuthenticationManagerProxyAdapter {
	return AuthenticationManagerProxyAdapter{
		ctx:          ctx,
		controlPlane: controlPlane,
	}
}

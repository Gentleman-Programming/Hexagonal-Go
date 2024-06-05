package adapters

import (
	control_plane "HEXAGONAL-GO/cmd/services/control-plane"
	"HEXAGONAL-GO/cmd/services/control-plane/models"
	"context"
)

type AuthenticationManagerProxyMockAdapter struct {
	ctx          context.Context
	controlPlane *control_plane.ControlPlane
}

func (a AuthenticationManagerProxyMockAdapter) Authenticate(username string, password string) (*models.User, error) {
	return a.controlPlane.Authenticate(a.ctx, username, password)
}

func NewAuthenticationManagerProxyMockAdapter(ctx context.Context, controlPlane *control_plane.ControlPlane) AuthenticationManagerProxyMockAdapter {
	return AuthenticationManagerProxyMockAdapter{
		ctx:          ctx,
		controlPlane: controlPlane,
	}
}

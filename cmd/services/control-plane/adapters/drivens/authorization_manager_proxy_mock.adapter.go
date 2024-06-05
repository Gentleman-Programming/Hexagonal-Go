package adapters

import (
	control_plane "HEXAGONAL-GO/cmd/services/control-plane"
	"context"
)

type AuthorizationManagerProxyAdapter struct {
	ctx          context.Context
	controlPlane *control_plane.ControlPlane
}

func (a *AuthorizationManagerProxyAdapter) Authorize(username string) (bool, error) {
	return a.controlPlane.Authorize(a.ctx, username)
}

func NewAuthorizationManagerProxyAdapter(ctx context.Context, controlPlane *control_plane.ControlPlane) AuthorizationManagerProxyAdapter {
	return AuthorizationManagerProxyAdapter{
		ctx:          ctx,
		controlPlane: controlPlane,
	}
}

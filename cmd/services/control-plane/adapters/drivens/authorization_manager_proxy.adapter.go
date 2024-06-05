package adapters

import (
	control_plane "HEXAGONAL-GO/cmd/services/control-plane"
	"context"
)

type AuthorizerProxyAdapter struct {
	ctx          context.Context
	controlPlane *control_plane.ControlPlane
}

func (a *AuthorizerProxyAdapter) Authorize(username string) (bool, error) {
	return a.controlPlane.Authorize(a.ctx, username)
}

func NewAuthorizerProxyAdapter(ctx context.Context, controlPlane *control_plane.ControlPlane) AuthorizerProxyAdapter {
	return AuthorizerProxyAdapter{
		ctx:          ctx,
		controlPlane: controlPlane,
	}
}

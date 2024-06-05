package adapters

import (
	control_plane "HEXAGONAL-GO/cmd/services/control-plane"
	"context"
)

type AuthorizerProxyAdapterMock struct {
	ctx          context.Context
	controlPlane *control_plane.ControlPlane
}

func (a *AuthorizerProxyAdapterMock) Authorize(username string) (bool, error) {
	return true, nil
}

func NewAuthorizerProxyAdapterMock(ctx context.Context, controlPlane *control_plane.ControlPlane) AuthorizerProxyAdapterMock {
	return AuthorizerProxyAdapterMock{
		ctx:          ctx,
		controlPlane: controlPlane,
	}
}

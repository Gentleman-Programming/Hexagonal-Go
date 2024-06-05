package control_plane

import (
	"HEXAGONAL-GO/cmd/services/control-plane/models"
	drivens "HEXAGONAL-GO/cmd/services/control-plane/ports/drivens"
	"context"
)

type ControlPlane struct {
	authorizationManager  drivens.ForManagingAuthorization
	authenticationManager drivens.ForManagingAuthentication
}

func (controlPlane *ControlPlane) Authenticate(ctx context.Context, username string, password string) (*models.User, error) {
	return controlPlane.authenticationManager.Authenticate(username, password)
}

func (controlPlane *ControlPlane) Authorize(ctx context.Context, username string) (bool, error) {
	return controlPlane.authorizationManager.Authorize(username)
}

// Control Plane
func NewControlPlaneService() *ControlPlane {
	return &ControlPlane{}
}

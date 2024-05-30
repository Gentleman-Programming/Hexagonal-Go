package control_plane

import (
	"HEXAGONAL-GO/cmd/services/control-plane/models"
	"context"
)

type ControlPlane struct {
	// drivens
}

func (controlPlane *ControlPlane) Authenticate(ctx context.Context, username string, password string) (*models.User, error) {
	return &models.User{
		Username: "alan",
		Password: "123456",
	}, nil
}

func (controlPlane *ControlPlane) Authorize(ctx context.Context, username string) (bool, error) {
	return true, nil
}

// Control Plane
func NewControlPlaneService() *ControlPlane {
	return &ControlPlane{}
}

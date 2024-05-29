package dashboard_api

import (
	"HEXAGONAL-GO/cmd/services/dashboard-api/models"
	ports "HEXAGONAL-GO/cmd/services/dashboard-api/ports/drivens"
	"context"
)

type DashboardApi struct {
	userQueryer ports.ForQueryingUser
}

func (d *DashboardApi) GetUser(ctx context.Context, username string) (*models.User, error) {
	return d.userQueryer.GetUser(username)
}

func NewDashboardService(userQueryer ports.ForQueryingUser) *DashboardApi {
	return &DashboardApi{userQueryer: userQueryer}
}

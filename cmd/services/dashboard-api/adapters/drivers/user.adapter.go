package adapters

import (
	dashboard_api "HEXAGONAL-GO/cmd/services/dashboard-api"
	"HEXAGONAL-GO/cmd/services/dashboard-api/models"
	"context"
)

type UserAdapter struct {
	ctx          context.Context
	dashboardApi *dashboard_api.DashboardApi
}

func (a *UserAdapter) GetUser(username string) (*models.User, error) {
	return a.dashboardApi.GetUser(a.ctx, username)
}

func CreateUserAdapter(ctx context.Context, dashboardApi *dashboard_api.DashboardApi) *UserAdapter {
	return &UserAdapter{
		ctx:          ctx,
		dashboardApi: dashboardApi,
	}
}

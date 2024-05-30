package dashboard_api

import (
	"HEXAGONAL-GO/cmd/services/dashboard-api/models"
	ports "HEXAGONAL-GO/cmd/services/dashboard-api/ports/drivens"
	"context"
)

type DashboardApi struct {
	userAuthenticator ports.ForAuthenticatingUser
	userAuthorizer    ports.ForAuthorizingUser
	userQueryer       ports.ForQueryingUser
}

func (d *DashboardApi) GetUser(ctx context.Context, username string) (*models.User, error) {
	return d.userQueryer.GetUser(username)
}

func (d *DashboardApi) Authenticate(ctx context.Context, username string, password string) (*models.User, error) {
	return d.userAuthenticator.Authenticate(username, password)
}

func (d *DashboardApi) Authorize(ctx context.Context, username string, password string) (bool, error) {
	return d.userAuthorizer.Authorize(username)
}

func NewDashboardService(userQueryer ports.ForQueryingUser, userAuthenticator ports.ForAuthenticatingUser, userAuthorizer ports.ForAuthorizingUser) *DashboardApi {
	return &DashboardApi{userQueryer: userQueryer, userAuthenticator: userAuthenticator, userAuthorizer: userAuthorizer}
}

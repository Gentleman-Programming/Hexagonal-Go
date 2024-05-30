package main

import (
	control_plane "HEXAGONAL-GO/cmd/services/control-plane"
	control_plane_adapters "HEXAGONAL-GO/cmd/services/control-plane/adapters/drivers"
	dashboard_api "HEXAGONAL-GO/cmd/services/dashboard-api"
	adapters "HEXAGONAL-GO/cmd/services/dashboard-api/adapters/drivens"
	dashboard_adapters "HEXAGONAL-GO/cmd/services/dashboard-api/adapters/drivers"
	dashboard_ports "HEXAGONAL-GO/cmd/services/dashboard-api/ports/drivers"
	"HEXAGONAL-GO/cmd/services/repository"
	repository_adapters "HEXAGONAL-GO/cmd/services/repository/adapters/drivers"
	"context"
)

func Compose() dashboard_ports.ForUser {
	// Create ctx (context)
	ctx := context.Background()

	// Create Services
	repository := repository.NewRepositoryService()
	controlPlane := control_plane.NewControlPlaneService()

	// Create repository drivers
	userManagerProxyAdapter := repository_adapters.NewUserManagerProxyAdapter(ctx, repository)

	// Create control plane drivers
	userAuthenticationManagerAdapter := control_plane_adapters.NewAuthenticationManagerProxyAdapter(ctx, controlPlane)
	userAuthorizationManagerAdapter := control_plane_adapters.NewAuthorizationManagerProxyAdapter(ctx, controlPlane)

	// Create dashboard api drivens
	userAuthenticatorAdapter := adapters.NewUserAuthenticatorAdapter(ctx, &userAuthenticationManagerAdapter)
	userAuthorizerAdapter := adapters.NewUserAuthorizerAdapter(ctx, &userAuthorizationManagerAdapter)
	userQueryerAdapter := adapters.NewUserQueryerAdapter(ctx, &userManagerProxyAdapter)

	// Create dashboard api
	dashboardApi := dashboard_api.NewDashboardService(userQueryerAdapter, userAuthenticatorAdapter, userAuthorizerAdapter)

	// Create dashboard api drivers
	userAdapter := dashboard_adapters.CreateUserAdapter(ctx, dashboardApi)

	return userAdapter
}

func ComposeMock() dashboard_ports.ForUser {
	// Create ctx (context)
	ctx := context.Background()

	// Create Services
	repository := repository.NewRepositoryService()
	controlPlane := control_plane.NewControlPlaneService()

	// Create repository drivers
	userManagerProxyAdapter := repository_adapters.NewUserManagerProxyAdapter(ctx, repository)

	// Create control plane drivens
	userAuthenticationManagerAdapter := control_plane_adapters.NewAuthenticationManagerProxyAdapter(ctx, controlPlane)
	userAuthorizationManagerAdapter := control_plane_adapters.NewAuthorizationManagerProxyAdapter(ctx, controlPlane)

	// Create dashboard api drivens
	userAuthenticator := adapters.NewUserAuthenticatorAdapter(ctx, &userAuthenticationManagerAdapter)
	userAuthorizer := adapters.NewUserAuthorizerAdapter(ctx, &userAuthorizationManagerAdapter)
	userQueryerAdapter := adapters.NewUserQueryerMockAdapter(ctx, &userManagerProxyAdapter)

	// Create dashboard api
	dashboardApi := dashboard_api.NewDashboardService(userQueryerAdapter, userAuthenticator, userAuthorizer)
	// Create dashboard api drivers
	userAdapter := dashboard_adapters.CreateUserAdapter(ctx, dashboardApi)

	return userAdapter
}

package main

import (
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

	// Create Repository
	repository := repository.NewRepository()

	// Create repository drivers
	userManagerProxyAdapter := repository_adapters.NewUserManagerProxyAdapter(ctx, repository)

	// Create dashboard api drivens
	userQueryerAdapter := adapters.NewUserQueryerAdapter(ctx, &userManagerProxyAdapter)

	// Create dashboard api
	dashboardApi := dashboard_api.NewDashboardService(userQueryerAdapter)

	// Create dashboard api drivers
	userAdapter := dashboard_adapters.CreateUserAdapter(ctx, dashboardApi)

	return userAdapter
}

func ComposeMock() dashboard_ports.ForUser {
	// Create ctx (context)
	ctx := context.Background()

	// Create Repository
	repository := repository.NewRepository()

	// Create repository drivers
	userManagerProxyAdapter := repository_adapters.NewUserManagerProxyAdapter(ctx, repository)

	// Create dashboard api drivens
	userQueryerAdapter := adapters.NewUserQueryerMockAdapter(ctx, &userManagerProxyAdapter)

	// Create dashboard api
	dashboardApi := dashboard_api.NewDashboardService(userQueryerAdapter)

	// Create dashboard api drivers
	userAdapter := dashboard_adapters.CreateUserAdapter(ctx, dashboardApi)

	return userAdapter
}

package adapters

import (
	"HEXAGONAL-GO/cmd/services/dashboard-api/models"
	repository_drivers "HEXAGONAL-GO/cmd/services/repository/ports/drivers"
	"context"
)

type UserQueryerMockAdapter struct {
	ctx    context.Context
	driver repository_drivers.ForManagingUser
}

func (a *UserQueryerMockAdapter) GetUser(username string) (*models.User, error) {
	mockedUser := models.User{
		Username: "alanbuscaglia",
		Password: "123456",
	}

	return &mockedUser, nil
}

func NewUserQueryerMockAdapter(ctx context.Context, driver repository_drivers.ForManagingUser) *UserQueryerMockAdapter {
	return &UserQueryerMockAdapter{
		ctx:    ctx,
		driver: driver,
	}
}

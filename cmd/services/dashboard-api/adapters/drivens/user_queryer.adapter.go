package adapters

import (
	"HEXAGONAL-GO/cmd/services/dashboard-api/models"
	repository_drivers "HEXAGONAL-GO/cmd/services/repository/ports/drivers"
	"context"
)

type UserQueryerAdapter struct {
	ctx    context.Context
	driver repository_drivers.ForManagingUser
}

func (a *UserQueryerAdapter) GetUser(username string) (*models.User, error) {
	user, err := a.driver.GetUser(username)
	if err != nil {
		return nil, err
	}

	mappedUser := models.User{
		Username: user.Username,
		Password: user.Password,
	}

	return &mappedUser, nil
}

func NewUserQueryerAdapter(ctx context.Context, driver repository_drivers.ForManagingUser) *UserQueryerAdapter {
	return &UserQueryerAdapter{
		ctx:    ctx,
		driver: driver,
	}
}

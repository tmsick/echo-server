package controller

import (
	"context"
	"echo-server/domain"

	"go.uber.org/zap"
)

type UsersController interface {
	ListUsers(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	RemoveUser(ctx context.Context, id string) error
}

type UsersControllerImpl struct {
	logger  func(ctx context.Context) *zap.Logger
	service domain.UsersAppService
}

func NewUsersControllerImpl(logger func(ctx context.Context) *zap.Logger, service domain.UsersAppService) *UsersControllerImpl {
	return &UsersControllerImpl{
		logger:  logger,
		service: service,
	}
}

func (c *UsersControllerImpl) ListUsers(ctx context.Context) ([]*User, error) {
	_users, err := c.service.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return FromUserDomainObjectSlice(_users), nil
}

func (c *UsersControllerImpl) GetUser(ctx context.Context, id string) (*User, error) {
	_user, err := c.service.GetUser(ctx, domain.UserID(id))
	if err != nil {
		return nil, err
	}
	return FromUserDomainObject(_user), nil
}

func (c *UsersControllerImpl) CreateUser(ctx context.Context, user *User) (*User, error) {
	_user, err := c.service.CreateUser(ctx, ToUserDomainObject(user))
	if err != nil {
		return nil, err
	}
	return FromUserDomainObject(_user), nil
}

func (c *UsersControllerImpl) UpdateUser(ctx context.Context, user *User) (*User, error) {
	_user, err := c.service.UpdateUser(ctx, ToUserDomainObject(user))
	if err != nil {
		return nil, err
	}
	return FromUserDomainObject(_user), nil
}

func (c *UsersControllerImpl) RemoveUser(ctx context.Context, id string) error {
	return c.service.RemoveUser(ctx, domain.UserID(id))
}

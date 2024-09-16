package domain

import (
	"context"
	"echo-server/repository"

	"go.uber.org/zap"
)

type UsersAppService interface {
	ListUsers(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id UserID) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	RemoveUser(ctx context.Context, id UserID) error
}

type UsersAppServiceImpl struct {
	logger     func(ctx context.Context) *zap.Logger
	repository repository.UsersRepository
}

func NewUsersAppServiceImpl(logger func(ctx context.Context) *zap.Logger, repository repository.UsersRepository) *UsersAppServiceImpl {
	return &UsersAppServiceImpl{
		logger:     logger,
		repository: repository,
	}
}

func (a *UsersAppServiceImpl) ListUsers(ctx context.Context) ([]*User, error) {
	_users, err := a.repository.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return FromUserRepositoryDTOSlice(_users), nil
}

func (a *UsersAppServiceImpl) GetUser(ctx context.Context, id UserID) (*User, error) {
	_user, err := a.repository.GetUser(ctx, id.String())
	if err != nil {
		return nil, err
	}
	return FromUserRepositoryDTO(_user), nil
}

func (a *UsersAppServiceImpl) CreateUser(ctx context.Context, user *User) (*User, error) {
	_user, err := a.repository.CreateUser(ctx, ToUserRepositoryDTO(user))
	if err != nil {
		return nil, err
	}
	return FromUserRepositoryDTO(_user), nil
}

func (a *UsersAppServiceImpl) UpdateUser(ctx context.Context, user *User) (*User, error) {
	_user, err := a.repository.UpdateUser(ctx, ToUserRepositoryDTO(user))
	if err != nil {
		return nil, err
	}
	return FromUserRepositoryDTO(_user), nil
}

func (a *UsersAppServiceImpl) RemoveUser(ctx context.Context, id UserID) error {
	return a.repository.RemoveUser(ctx, id.String())
}

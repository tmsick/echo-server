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
	users, err := a.repository.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*User, 0, len(users))
	for _, u := range users {
		list = append(list,
			&User{
				ID:    UserID(u.ID),
				Name:  UserName(u.Name),
				Email: UserEmail(u.Email),
			})
	}
	return list, nil
}

func (a *UsersAppServiceImpl) GetUser(ctx context.Context, id UserID) (*User, error) {
	user, err := a.repository.GetUser(ctx, id.String())
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    UserID(user.ID),
		Name:  UserName(user.Name),
		Email: UserEmail(user.Email),
	}, nil
}

func (a *UsersAppServiceImpl) CreateUser(ctx context.Context, u *User) (*User, error) {
	user, err := a.repository.CreateUser(ctx, &repository.User{
		ID:    u.ID.String(),
		Name:  u.Name.String(),
		Email: u.Email.String(),
	})
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    UserID(user.ID),
		Name:  UserName(user.Name),
		Email: UserEmail(user.Email),
	}, nil
}

func (a *UsersAppServiceImpl) UpdateUser(ctx context.Context, u *User) (*User, error) {
	user, err := a.repository.UpdateUser(ctx, &repository.User{
		ID:    u.ID.String(),
		Name:  u.Name.String(),
		Email: u.Email.String(),
	})
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    UserID(user.ID),
		Name:  UserName(user.Name),
		Email: UserEmail(user.Email),
	}, nil
}

func (a *UsersAppServiceImpl) RemoveUser(ctx context.Context, id UserID) error {
	return a.repository.RemoveUser(ctx, id.String())
}

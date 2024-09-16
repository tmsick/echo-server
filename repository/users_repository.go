package repository

import (
	"context"
	"errors"
	"strconv"

	"go.uber.org/zap"
)

var Users = map[string]*User{
	"1": {
		ID:       "1",
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "password_alice",
	},
	"2": {
		ID:       "2",
		Name:     "Bob",
		Email:    "bob@example.com",
		Password: "password_bob",
	},
}

var ErrUserNotFound = errors.New("user not found")

type UsersRepository interface {
	ListUsers(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	RemoveUser(ctx context.Context, id string) error
}

type UsersRepositoryImpl struct {
	logger func(ctx context.Context) *zap.Logger
}

func NewUsersRepositoryImpl(logger func(ctx context.Context) *zap.Logger) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{
		logger: logger,
	}
}

func (r *UsersRepositoryImpl) ListUsers(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0, len(Users))
	for _, user := range Users {
		users = append(users, user)
	}
	return users, nil
}

func (r *UsersRepositoryImpl) GetUser(ctx context.Context, id string) (*User, error) {
	user, ok := Users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *UsersRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	for _, user := range Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (r *UsersRepositoryImpl) CreateUser(ctx context.Context, user *User) (*User, error) {
	id := strconv.Itoa(len(Users) + 1)
	user.ID = id
	Users[id] = user
	return user, nil
}

func (r *UsersRepositoryImpl) UpdateUser(ctx context.Context, user *User) (*User, error) {
	_, ok := Users[user.ID]
	if !ok {
		return nil, errors.New("user not found")
	}
	Users[user.ID] = user
	return user, nil
}

func (r UsersRepositoryImpl) RemoveUser(ctx context.Context, id string) error {
	_, ok := Users[id]
	if !ok {
		return errors.New("user not found")
	}
	delete(Users, id)
	return nil
}

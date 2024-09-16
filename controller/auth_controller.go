package controller

import (
	"context"
	"errors"

	"github.com/tmsick/echo-server/domain"
	"go.uber.org/zap"
)

var (
	ErrInvalidCredential = errors.New("invalid credential")
	ErrUserNotFound      = errors.New("user not found")
)

type AuthController interface {
	SignIn(ctx context.Context, credential *SignInCredential) (*AuthToken, error)
}

type AuthControllerImpl struct {
	logger  func(ctx context.Context) *zap.Logger
	service domain.AuthAppService
}

func NewAuthControllerImpl(logger func(ctx context.Context) *zap.Logger, service domain.AuthAppService) *AuthControllerImpl {
	return &AuthControllerImpl{
		logger:  logger,
		service: service,
	}
}

func (c *AuthControllerImpl) SignIn(ctx context.Context, credential *SignInCredential) (*AuthToken, error) {
	_token, err := c.service.SignIn(ctx, &domain.SignInCredential{
		Email:    domain.UserEmail(credential.Email),
		Password: domain.UserPassword(credential.Password),
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidCredential):
			return nil, ErrInvalidCredential
		case errors.Is(err, domain.ErrUserNotFound):
			return nil, ErrUserNotFound
		default:
			return nil, err
		}
	}
	return &AuthToken{
		Token: _token.Token,
	}, nil
}

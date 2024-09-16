package domain

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tmsick/echo-server/environment"
	"github.com/tmsick/echo-server/kontext"
	"github.com/tmsick/echo-server/repository"
	"go.uber.org/zap"
)

var (
	ErrInvalidCredential = errors.New("invalid credential")
	ErrUserNotFound      = errors.New("user not found")
)

type AuthAppService interface {
	SignIn(ctx context.Context, credential *SignInCredential) (*AuthToken, error)
}

type AuthAppServiceImpl struct {
	logger func(ctx context.Context) *zap.Logger
	users  repository.UsersRepository
}

func NewAuthAppServiceImpl(
	logger func(ctx context.Context) *zap.Logger,
	users repository.UsersRepository,
) *AuthAppServiceImpl {
	return &AuthAppServiceImpl{
		logger: logger,
		users:  users,
	}
}

func (a *AuthAppServiceImpl) SignIn(ctx context.Context, credential *SignInCredential) (*AuthToken, error) {
	user, err := a.users.GetUserByEmail(ctx, credential.Email.String())
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if credential.Password.String() != user.Password {
		return nil, ErrInvalidCredential
	}

	issuedAt := kontext.GetRequestTime(ctx)
	expiresAt := issuedAt.Add(24 * time.Hour)
	jti := uuid.New().String()

	claims := jwt.RegisteredClaims{
		Issuer:    "echo-server",
		Subject:   credential.Email.String(),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ID:        jti,
	}

	logger := a.logger(ctx)
	defer logger.Sync()
	logger.Info("jwt created",
		zap.Dict("jwt",
			zap.String("id", claims.ID),
			zap.String("issuer", claims.Issuer),
			zap.String("subject", claims.Subject),
			zap.Time("expires_at", claims.ExpiresAt.Time),
			zap.Time("issued_at", claims.IssuedAt.Time),
		),
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(environment.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		Token: ss,
	}, nil
}

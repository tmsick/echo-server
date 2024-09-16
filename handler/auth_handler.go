package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmsick/echo-server/controller"
	"go.uber.org/zap"
)

var (
	ErrInvalidCredential = errors.New("invalid credential")
	ErrUserNotFound      = errors.New("user not found")
)

type AuthHandler interface {
	Register(g *echo.Group)
	Create(c echo.Context) error
}

type AuthHandlerImpl struct {
	logger     func(ctx context.Context) *zap.Logger
	controller controller.AuthController
}

func NewAuthHandlerImpl(logger func(ctx context.Context) *zap.Logger, controller controller.AuthController) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		logger:     logger,
		controller: controller,
	}
}

func (h *AuthHandlerImpl) Register(g *echo.Group) {
	g.POST("/sign-in", h.Create)
}

func (h *AuthHandlerImpl) Create(c echo.Context) error {
	ctx := c.Request().Context()

	credential := new(SignInCredential)
	if err := c.Bind(credential); err != nil {
		return err
	}
	if err := c.Validate(credential); err != nil {
		return err
	}

	token, err := h.controller.SignIn(ctx, &controller.SignInCredential{
		Email:    credential.Email,
		Password: credential.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, controller.ErrInvalidCredential):
			return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidCredential)
		case errors.Is(err, controller.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, ErrUserNotFound)
		default:
			return err
		}
	}

	return c.JSON(200, &AuthToken{
		Token: token.Token,
	})
}

package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmsick/echo-server/controller"
	"go.uber.org/zap"
)

type UsersHandler interface {
	Register(g *echo.Group)
	Index(c echo.Context) error
	Show(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Remove(c echo.Context) error
}

type UsersHandlerImpl struct {
	logger     func(ctx context.Context) *zap.Logger
	controller controller.UsersController
}

func NewUsersHandlerImpl(logger func(ctx context.Context) *zap.Logger, controller controller.UsersController) *UsersHandlerImpl {
	return &UsersHandlerImpl{
		logger:     logger,
		controller: controller,
	}
}

func (h *UsersHandlerImpl) Register(g *echo.Group) {
	g.GET("", h.Index)
	g.GET("/:id", h.Show)
	g.POST("", h.Create)
	g.PATCH("/:id", h.Update)
	g.DELETE("/:id", h.Remove)
}

func (h *UsersHandlerImpl) Index(c echo.Context) error {
	ctx := c.Request().Context()
	_users, err := h.controller.ListUsers(ctx)
	if err != nil {
		return err
	}
	users := FromUserControllerDTOSlice(_users)
	return c.JSON(http.StatusOK, users)
}

func (h *UsersHandlerImpl) Show(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	_user, err := h.controller.GetUser(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(200, FromUserControllerDTO(_user))
}

func (h *UsersHandlerImpl) Create(c echo.Context) error {
	ctx := c.Request().Context()

	logger := h.logger(ctx)
	defer logger.Sync()

	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}

	_user, err := h.controller.CreateUser(ctx, ToUserControllerDTO(user))
	if err != nil {
		return err
	}
	return c.JSON(200, FromUserControllerDTO(_user))
}

func (h *UsersHandlerImpl) Update(c echo.Context) error {
	ctx := c.Request().Context()
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	_user, err := h.controller.UpdateUser(ctx, ToUserControllerDTO(user))
	if err != nil {
		return err
	}
	return c.JSON(200, FromUserControllerDTO(_user))
}

func (h *UsersHandlerImpl) Remove(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	err := h.controller.RemoveUser(ctx, id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

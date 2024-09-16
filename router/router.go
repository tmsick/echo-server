package router

import (
	"echo-server/controller"
	"echo-server/domain"
	"echo-server/handler"
	"echo-server/logger"
	"echo-server/repository"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Register(e *echo.Echo, l *zap.Logger) {
	// ---- users ----
	{
		var r repository.UsersRepository = repository.NewUsersRepositoryImpl(logger.WithContext(l))
		var s domain.UsersAppService = domain.NewUsersAppServiceImpl(logger.WithContext(l), r)
		var c controller.UsersController = controller.NewUsersControllerImpl(logger.WithContext(l), s)
		var h handler.UsersHandler = handler.NewUsersHandlerImpl(logger.WithContext(l), c)
		g := e.Group("/users")
		g.GET("", h.Index)
		g.GET("/:id", h.Show)
		g.POST("", h.Create)
		g.PATCH("/:id", h.Update)
		g.DELETE("/:id", h.Remove)
	}
}

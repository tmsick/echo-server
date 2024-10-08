package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tmsick/echo-server/controller"
	"github.com/tmsick/echo-server/domain"
	"github.com/tmsick/echo-server/environment"
	"github.com/tmsick/echo-server/handler"
	"github.com/tmsick/echo-server/kontext"
	"github.com/tmsick/echo-server/logger"
	"github.com/tmsick/echo-server/repository"
	"github.com/tmsick/echo-server/validator"
	"go.uber.org/zap"
)

func main() {
	// ---- logger ----
	l, err := logger.New(environment.Env)
	if err != nil {
		panic(err)
	}
	defer l.Sync()

	// ---- echo ----
	e := echo.New()

	// ---- validator ----
	e.Validator = validator.New()

	// ---- middleware ----
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			now := time.Now()
			ctx = kontext.SetRequestTime(ctx, now)
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			id := uuid.New().String()
			ctx = kontext.SetRequestID(ctx, id)
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t0 := time.Now()

			req := c.Request()
			res := c.Response()

			ctx := req.Context()
			id := kontext.GetRequestID(ctx)
			reqTime := kontext.GetRequestTime(ctx)
			l.Info("api_start",
				zap.String("host", req.Host),
				zap.String("method", req.Method),
				zap.String("referer", req.Referer()),
				zap.String("remote_ip", c.RealIP()),
				zap.String("request_id", id),
				zap.String("uri", req.RequestURI),
				zap.String("user_agent", req.UserAgent()),
				zap.Time("request_time", reqTime),
			)

			err := next(c)
			c.Error(err)

			t1 := time.Now()
			td := t1.Sub(t0)

			id = kontext.GetRequestID(req.Context())
			reqTime = kontext.GetRequestTime(ctx)
			l.Info("api_end",
				zap.Error(err),
				zap.Int("status", res.Status),
				zap.Int64("latency_nano", int64(td)),
				zap.Int64("response_size", res.Size),
				zap.String("request_id", id),
				zap.Stringer("latency_human", td),
				zap.Time("request_time", reqTime),
			)
			return err
		}
	})

	// ---- DI ----
	var (
		// ---- users ----
		usersRepository repository.UsersRepository = repository.NewUsersRepositoryImpl(logger.WithContext(l))
		usersAppService domain.UsersAppService     = domain.NewUsersAppServiceImpl(logger.WithContext(l), usersRepository)
		usersController controller.UsersController = controller.NewUsersControllerImpl(logger.WithContext(l), usersAppService)
		usersHandler    handler.UsersHandler       = handler.NewUsersHandlerImpl(logger.WithContext(l), usersController)

		// ---- auth ----
		authAppService domain.AuthAppService     = domain.NewAuthAppServiceImpl(logger.WithContext(l), usersRepository)
		authController controller.AuthController = controller.NewAuthControllerImpl(logger.WithContext(l), authAppService)
		authHandler    handler.AuthHandler       = handler.NewAuthHandlerImpl(logger.WithContext(l), authController)
	)

	// ---- register routers ----
	usersHandler.Register(e.Group("/users"))
	authHandler.Register(e.Group("/auth"))

	// ---- start ----
	l.Fatal("fatal server", zap.Error(e.Start("127.0.0.1:8080")))
}

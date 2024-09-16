package main

import (
	"echo-server/controller"
	"echo-server/domain"
	"echo-server/environment"
	"echo-server/handler"
	"echo-server/kontext"
	"echo-server/logger"
	"echo-server/repository"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	// ---- env ----
	env, err := environment.Load()
	if err != nil {
		panic(err)
	}

	// ---- logger ----
	l, err := logger.New(env.Env)
	if err != nil {
		panic(err)
	}
	defer l.Sync()

	// ---- echo ----
	e := echo.New()

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

			t1 := time.Now()
			td := t1.Sub(t0)

			req = c.Request()
			res := c.Response()
			id = kontext.GetRequestID(req.Context())
			reqTime = kontext.GetRequestTime(ctx)
			l.Info("api_end",
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
		usersRepository repository.UsersRepository = repository.NewUsersRepositoryImpl(logger.WithContext(l))
		usersAppService domain.UsersAppService     = domain.NewUsersAppServiceImpl(logger.WithContext(l), usersRepository)
		usersController controller.UsersController = controller.NewUsersControllerImpl(logger.WithContext(l), usersAppService)
		usersHandler    handler.UsersHandler       = handler.NewUsersHandlerImpl(logger.WithContext(l), usersController)
	)

	// ---- register routers ----
	usersHandler.Register(e.Group("/users"))

	// ---- start ----
	l.Fatal("fatal server", zap.Error(e.Start("127.0.0.1:8080")))
}

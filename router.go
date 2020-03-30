package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

type Option func(*echo.Echo)

var ops = make([]Option, 0, 10)

func WarpEcho(e *echo.Echo, os ...Option) {
	for _, v := range os {
		v(e)
	}
}

func initRouter() {
	defer func() {
		e.Close()
	}()
	e = echo.New()
	e.Use(ZeroLoggerWithConfig(LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           logger,
	}), middleware.Recover(), middleware.CORS(), middleware.Gzip())
	e.HideBanner = true
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	WarpEcho(e, ops...)
	logger.Info().Msg("已注册路由列表：")
	for _, v := range e.Routes() {
		logger.Info().Str("router.mother", v.Method).Str("router.path", v.Path).Msg("")
	}
	go func() {
		if err := e.Start(conf.Port); err != nil {
			logger.Error().Msg("start server error:" + err.Error())
			logger.Fatal().Msg("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal().Msg(err.Error())
	}
}

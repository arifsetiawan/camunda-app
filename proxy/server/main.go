package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"gopkg.in/go-playground/validator.v9"

	"github.com/arifsetiawan/camunda-app/pkg/env"
	pkgMiddleware "github.com/arifsetiawan/camunda-app/pkg/middleware"
	"github.com/arifsetiawan/camunda-app/proxy/handler"
)

const appName = "camunda-app"

type (
	// CustomValidator is
	CustomValidator struct {
		validator *validator.Validate
	}
)

// Validate is
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	// @
	// Initialize echo
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(pkgMiddleware.Logger(appName))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	p := pkgMiddleware.NewPrometheus("camunda-app")
	p.Use(e)

	handler.InitHandler(e)

	// @
	// Start app
	servicePort := env.Getenv("PORT", ":9010")
	log.Info().Msg(appName + " with version: " + env.Getenv("VERSION", "0.1.0") + " started at port " + servicePort)

	go func() {
		if err := e.Start(servicePort); err != nil {
			log.Info().Err(err).Msg(appName + " is shutting down")
		}
	}()

	// @
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout
	shutdownTimeout, err := time.ParseDuration(env.Getenv("SHUTDOWN_TIMEOUT", "1s"))
	if err != nil {
		log.Info().Err(err).Msg("Failed to parse shutdown timeout")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	s := <-quit
	log.Info().Err(err).Interface("signal", s).Msg("Got signal")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

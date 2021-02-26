package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type (
	// LoggerConfig defines the config for Logger middleware.
	LoggerConfig struct {
		Skipper Skipper
		AppName string
	}
)

// DefaultLoggerConfig ...
func DefaultLoggerConfig(appName string) LoggerConfig {
	return LoggerConfig{
		Skipper: DefaultSkipper,
		AppName: appName,
	}
}

// Logger returns a middleware that logs HTTP requests.
func Logger(appName string) echo.MiddlewareFunc {
	return LoggerWithConfig(DefaultLoggerConfig(appName))
}

// LoggerWithConfig returns a Logger middleware with config.
// See: `Logger()`.
func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			tenant := c.Param("tenant")
			if tenant == "" {
				tenant = "none"
			}

			path := req.URL.Path
			if path == "" {
				path = "/"
			}

			latency := stop.Sub(start).Nanoseconds() / int64(time.Microsecond)
			latencyHuman := stop.Sub(start).String()
			bytesIn := req.Header.Get(echo.HeaderContentLength)
			if bytesIn == "" {
				bytesIn = "0"
			}

			log.Info().
				Str("tenant", tenant).
				Str("type", "request").
				Str("remote_ip", c.RealIP()).
				Str("host", req.Host).
				Str("uri", req.RequestURI).
				Str("path", path).
				Str("method", req.Method).
				Str("referer", req.Referer()).
				Str("user_agent", req.UserAgent()).
				Int("status", res.Status).
				Int64("latency", latency).
				Str("latency_human", latencyHuman).
				Str("bytes_in", bytesIn).
				Str("bytes_out", strconv.FormatInt(res.Size, 10)).
				Msg("Request handled")

			// Todo: add body and header to log

			return nil
		}
	}
}

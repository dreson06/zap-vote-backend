package simplelog

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"time"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		start := time.Now()
		err := next(e)
		log.Debug().
			CallerSkipFrame(10).
			Str("\t"+e.Request().Method, e.Path()).
			Str("t", time.Since(start).String()).
			Send()
		return err
	}
}

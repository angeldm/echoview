package middleware

import (
	angeldm "angeldm.echoview/application/context"
	"github.com/labstack/echo/v4"

	"upper.io/db.v3/sqlite"
)

var settings = sqlite.ConnectionURL{
	Database: "echoview_development.sqlite",
}

// NewMiddleware echo middleware for func `echoview.Render()`
func UpperdbPOPMiddleware() echo.MiddlewareFunc {
	return UpperdbMiddleware()
}

// Middleware echo middleware wrapper
func UpperdbMiddleware() echo.MiddlewareFunc {
	sess, err := sqlite.Open(settings)
	if err != nil {
		panic(err)
	}
	sess.SetLogging(true)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &angeldm.CustomContext{Context: c, DB: sess}
			return next(cc)
		}
	}
}

package middleware

//
//import (
//	"angeldm.echoview/application/ctx"
//	"angeldm.echoview/utils/logmatic"
//	"github.com/gobuffalo/pop"
//	"github.com/gobuffalo/pop/logging"
//	"github.com/labstack/echo/v4"
//)
//
//var defaultLogger = func(lvl logging.Level, s string, args ...interface{}) {
//	l := logmatic.NewLogger("POP")
//	l.SetLevel(logmatic.DEBUG)
//	l.Log(lvl.String(), s, args)
//}
//
//// NewMiddleware echo middleware for func `echoview.Render()`
//func NewPOPMiddleware() echo.MiddlewareFunc {
//	return POPMiddleware()
//}
//
//// Middleware echo middleware wrapper
//func POPMiddleware() echo.MiddlewareFunc {
//	pop.SetLogger(defaultLogger)
//	_, err := pop.Connect("development")
//
//	if err != nil {
//		panic(err)
//	}
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			cc := &ctx.CustomContext{Context: c}
//			return next(cc)
//		}
//	}
//}

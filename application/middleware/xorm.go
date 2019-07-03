package middleware

//import (
//	angeldm "angeldm.echoview/application/ctx"
//	"github.com/go-xorm/xorm"
//	"github.com/labstack/echo/v4"
//	_ "github.com/mattn/go-sqlite3"
//	"os"
//)
//
//// NewMiddleware echo middleware for func `echoview.Render()`
//func NewXormMiddleware() echo.MiddlewareFunc {
//	return XormMiddleware()
//}
//
//// Middleware echo middleware wrapper
//func XormMiddleware() echo.MiddlewareFunc {
//
//	engine, err := xorm.NewEngine("sqlite3", "echoview.db")
//	logger := xorm.NewSimpleLogger(os.Stdout)
//	logger.ShowSQL(true)
//	engine.SetLogger(logger)
//	if err != nil {
//		panic(err)
//	}
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			cont := &angeldm.Context{Context: c, Engine: engine}
//			return next(cont)
//		}
//	}
//}

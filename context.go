package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/logging"
	"github.com/labstack/echo/v4"
	stdlog "log"
	"os"
)

type logger func(lvl logging.Level, s string, args ...interface{})

var defaultStdLogger = stdlog.New(os.Stdout, "[POP] ", stdlog.LstdFlags)

// Debug mode, to toggle verbose log traces
var Debug = true

// Color mode, to toggle colored logs
var Color = true

var defaultLogger = func(lvl logging.Level, s string, args ...interface{}) {
	// Handle legacy logger
	if !Debug && lvl <= logging.Debug {
		return
	}
	if lvl == logging.SQL {
		if len(args) > 0 {
			xargs := make([]string, len(args))
			for i, a := range args {
				switch a.(type) {
				case string:
					xargs[i] = fmt.Sprintf("%q", a)
				default:
					xargs[i] = fmt.Sprintf("%v", a)
				}
			}
			s = fmt.Sprintf("%s - %s | %s", lvl, s, xargs)
		} else {
			s = fmt.Sprintf("%s - %s", lvl, s)
		}
	} else {
		s = fmt.Sprintf(s, args...)
		s = fmt.Sprintf("%s - %s", lvl, s)
	}
	if Color {
		s = color.YellowString(s)
	}
	defaultStdLogger.Println(s)
}

type CustomContext struct {
	echo.Context
	Connection *pop.Connection
	//Log *logger
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

// NewMiddleware echo middleware for func `echoview.Render()`
func NewPOPMiddleware() echo.MiddlewareFunc {
	return POPMiddleware()
}

// Middleware echo middleware wrapper
func POPMiddleware() echo.MiddlewareFunc {
	pop.SetLogger(defaultLogger)
	conn, err := pop.Connect("development")

	if err != nil {
		panic(err)
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{Context: c, Connection: conn}
			return next(cc)
		}
	}
}

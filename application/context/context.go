package context

import (
	"github.com/gobuffalo/pop"
	"github.com/labstack/echo/v4"
)

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

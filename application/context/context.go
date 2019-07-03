package context

import (
	"github.com/labstack/echo/v4"
	"upper.io/db.v3/lib/sqlbuilder"
)

type CustomContext struct {
	echo.Context
	DB sqlbuilder.Database
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"angeldm.echoview/models"
	_ "github.com/mattn/go-sqlite3"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//conn, err := pop.Connect("development")
	//if err!= nil {
	//	panic(err)
	//}
	//users := models.Users{}
	//err = conn.All(&users)
	//if err!=nil {
	//	panic(err)
	//}else {
	//	fmt.Println(users)
	//}

	// Echo instance
	e := echo.New()

	e.Static("/static", "public/webpack")
	// Middleware
	//	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(NewPOPMiddleware())

	//Set Renderer
	e.Renderer = Default()

	// Routes
	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.Foo()
		users := models.Users{}
		err := cc.Connection.All(&users)
		if err != nil {
			panic(err)
		}
		//render 	with master
		return c.Render(http.StatusOK, "index", echo.Map{
			"title": "Index title!",
			"users": users,
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}

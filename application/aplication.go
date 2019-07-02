package application

import (
	"angeldm.echoview/application/context"
	angeldm "angeldm.echoview/application/middleware"
	"angeldm.echoview/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Application struct {
	*echo.Echo
}

func NewApplication() *Application {
	// Echo instance
	e := echo.New()
	e.Debug = true
	e.Static("/static", "public/webpack")
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] DEBUG method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(angeldm.NewPOPMiddleware())

	//Set Renderer
	e.Renderer = angeldm.Default()

	// Routes
	e.GET("/", func(c echo.Context) error {
		cc := c.(*context.CustomContext)
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

	return &Application{Echo: e}

}

func (a *Application) Start() {
	// Start server
	a.Echo.Logger.Fatal(a.Echo.Start(":9090"))
}

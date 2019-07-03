package server

import (
	"angeldm.echoview/application/controller/users"
	"angeldm.echoview/application/ctx"
	middleware2 "angeldm.echoview/application/middleware"
	"angeldm.echoview/utils/logger"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server is an main application object that shared (read-only) to application modules
type Server struct {
	context *ctx.Context
	echoSrv *echo.Echo
}

// New constructor
func New(c *ctx.Context) *Server {
	s := new(Server)
	s.context = c
	s.echoSrv = echo.New()
	return s
}

// Run registers API and starts http-server
func (s *Server) Run() {
	e := s.echoSrv
	e.Static("/static", "public/webpack")
	// Global Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(logger.HTTPLogger(s.context.Logger))
	e.Renderer = middleware2.Default()

	var (
		//	authHandler    = auth.NewHandler(s.context)
		//	versionHandler = version.NewHandler(s.context)
		//
		//)
		usersHandler = users.NewHandler(s.context)
	)
	// Non-authored routes
	//e.POST("/auth", authHandler.PostAuth)
	//e.GET("/", versionHandler.GetVersion)
	//e.GET("/version", versionHandler.GetVersion)
	//// restricted
	//r := e.Group("")
	//// group middleware
	//r.Use(middleware.JWT(s.context.JWTSignKey))
	//// users
	//r.POST("/users", usersHandler.PostUser)
	//r.PUT("/users/:id", usersHandler.PutUser)
	//r.GET("/users", usersHandler.GetAllUsers)
	//r.GET("/users/:id", usersHandler.GetUser)
	//r.DELETE("/user/:id", usersHandler.DeleteUser)
	e.GET("/", usersHandler.Index)
	// start server
	e.HideBanner = false
	e.Server.Addr = ":" + s.context.Config.Port
	s.context.Logger.WithField("addres", e.Server.Addr).Info("ECHO[SERVER] Starting server")
	err := e.Start(e.Server.Addr)
	if err != nil {
		s.context.Logger.Error("appcontrol", err.Error())
	}
}

// Shutdown gracefully stops server
func (s Server) Shutdown() error {
	return s.echoSrv.Shutdown(context.Background())
}

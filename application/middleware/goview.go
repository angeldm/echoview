package middleware

import (
	"angeldm.echoview/utils/goview"
	"angeldm.echoview/utils/logger"
	"github.com/go-webpack/webpack"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"html/template"
	"io"
)

const templateEngineKey = "foolin-goview-echoview"

// DefaultConfig default config
var DefaultConfig = goview.Config{
	Root:         "views",
	Extension:    ".html",
	Master:       "layouts/master",
	Partials:     []string{},
	Funcs:        make(template.FuncMap),
	DisableCache: true,
	Delims:       goview.Delims{Left: "{{", Right: "}}"},
}

// ViewEngine view engine for echo
type ViewEngine struct {
	echo.Context
	Logger *logrus.Logger
	*goview.ViewEngine
}

// New new view engine
func New(config goview.Config) *ViewEngine {
	webpack.Plugin = "manifest"
	webpack.Init(true)
	config.Funcs["asset"] = webpack.AssetHelper
	gv := goview.New(config)
	//l := logger.NewStdLogger("GOVIEW","")
	l := logger.NewLogrus()
	return &ViewEngine{

		Logger:     l,
		ViewEngine: gv,
	}
}

// Default new default config view engine
func Default() *ViewEngine {

	return New(DefaultConfig)
}

// Render render template for echo interface
func (e *ViewEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	fields := make(logrus.Fields)
	mapp := data.(echo.Map)
	for key, value := range mapp {
		fields[key] = value
	}
	e.Logger.WithFields(fields).Info("GOVIEW[" + name + "]")
	err := e.RenderWriter(w, name, data)
	if err != nil {
		e.Logger.Info(name)
	}
	return err
}

// Render html render for template
// You should use helper func `Middleware()` to set the supplied
// TemplateEngine and make `Render()` work validly.
func Render(ctx echo.Context, code int, name string, data interface{}) error {
	if val := ctx.Get(templateEngineKey); val != nil {
		if e, ok := val.(*ViewEngine); ok {
			return e.Render(ctx.Response().Writer, name, data, ctx)
		}
	}
	return ctx.Render(code, name, data)
}

// NewMiddleware echo middleware for func `echoview.Render()`
func NewMiddleware(config goview.Config) echo.MiddlewareFunc {
	return Middleware(New(config))
}

// Middleware echo middleware wrapper
func Middleware(e *ViewEngine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(templateEngineKey, e)
			return next(c)
		}
	}
}

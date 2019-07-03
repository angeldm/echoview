package application

import (
	"angeldm.echoview/application/ctx"
	"angeldm.echoview/application/server"
	"angeldm.echoview/models"
	"angeldm.echoview/utils/errors"
	"angeldm.echoview/utils/logger"
	"github.com/BurntSushi/toml"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Application define a mode of running app
type Application struct {
	Ctx *ctx.Context
	srv *server.Server
}

// New constructor
func New(flags *ctx.Flags) (*Application, error) {
	a := new(Application)
	a.Ctx = new(ctx.Context)
	// read config file
	err := a.initConfigFromFile(flags.CfgFileName)
	if err != nil {
		return nil, err
	}

	// init Logger
	a.initLogger()

	// connect to Db
	a.Ctx.Logger.Info("ECHO[APP] Started connection to database")
	err = a.initOrm()
	if err != nil {
		return nil, err
	}
	a.Ctx.Logger.Info("ECHO[APP] Connected to database successfully")

	a.Ctx.Config.Version = "0.1.0-dev"

	return a, nil
}

// Run starts application
func (a *Application) Run() {
	a.srv = server.New(a.Ctx)
	a.srv.Run()
}

// Shutdown gracefully stops server
func (a Application) Shutdown() error {
	// stop server
	a.Ctx.Logger.Info("ECHO[APP] Stopping server")
	if a.srv != nil {
		err := a.srv.Shutdown()
		if err != nil {
			return err
		}
	}

	// close database connection
	if a.Ctx.Orm != nil {
		a.Ctx.Logger.Info("ECHO[APP] Closing db connection")
		err := a.Ctx.Orm.Close()
		if err != nil {
			return err
		}
	}

	// close logger
	if a.Ctx.Logger != nil {
		a.Ctx.Logger.Info("ECHO[APP] Logger stopped")
		a.Ctx.Logger.Info("ECHO[APP] Quitting")
		//a.Ctx.Logger.Close()
	}
	return nil
}

//-----------------------------------------------------------------------------

// readConfig reads configuration file into application Config structure and inits in-memory token storage
func (a *Application) initConfigFromFile(cfgFileName string) error {
	// read config
	tomlData, err := ioutil.ReadFile(cfgFileName) //nolint
	if err != nil {
		return errors.New("Configuration file read error: " + cfgFileName + "\nError:" + err.Error())
	}
	_, err = toml.Decode(string(tomlData), &a.Ctx.Config)
	if err != nil {
		return errors.New("Configuration file decoding error: " + cfgFileName + "\nError:" + err.Error())
	}
	// init Logging data
	if a.Ctx.Config.Logging.ID == "" {
		a.Ctx.Config.Logging.ID = strconv.Itoa(os.Getpid())
	}
	if a.Ctx.Config.Logging.LogTag == "" {
		a.Ctx.Config.Logging.LogTag = os.Args[0]
	}
	return nil
}

// setupLogger sets apllication Logger up according to configuration settings
func (a *Application) initLogger() {
	//if a.Ctx.Config.Logging.LogMode == "nil" || a.Ctx.Config.Logging.LogMode == "null" {
	//	a.Ctx.Logger = logger.NewNilLogger()
	//	return
	//}
	a.Ctx.Logger = logger.NewLogrus() //logger.NewStdLogger(a.Ctx.Config.Logging.ID, a.Ctx.Config.Logging.LogTag)
}

// init database
func (a *Application) initOrm() error {
	var err error
	// open database
	a.Ctx.Orm, err = xorm.NewEngine(a.Ctx.Config.Database.Db, a.Ctx.Config.Database.Dsn)
	if err != nil {
		return err
	}

	// turn on logs

	ormLogger := logger.NewOrmLogger(logger.NewLogrus())
	a.Ctx.Orm.SetLogger(ormLogger)
	a.Ctx.Orm.SetLogLevel(core.LOG_DEBUG)
	a.Ctx.Orm.ShowSQL(true)
	// migrate
	err = a.migrateDb()
	if err != nil {
		return err
	}
	// init data
	err = a.initDbData()
	return err
}

// migrate database
func (a *Application) migrateDb() error {
	// migrate tables
	return a.Ctx.Orm.Sync(&models.User{})
}

// initDbData installs hardcoded data from config
func (a *Application) initDbData() error {
	user := &models.User{Email: "admin", DisplayName: "admin", Password: "admin"}      // aaaa, backdoor
	user1 := &models.User{Email: "angeldm", DisplayName: "angeldm", Password: "admin"} // aaaa, backdoor
	err := user.Save(a.Ctx.Orm)
	if err == nil {
		return nil
	}
	err = user1.Save(a.Ctx.Orm)
	if err == nil {
		return nil
	}
	status, _ := errors.Decompose(err)
	if status == http.StatusConflict {
		return nil
	}
	err = errors.NewWithPrefix(err, "database error")
	a.Ctx.Logger.Error("application init error", err.Error())
	return err
}

//type Application struct {
//	*echo.Echo
//}
//
//func NewApplication() *Application {
//	// Echo instance
//	e := echo.New()
//	e.Debug = true
//	e.Static("/static", "public/webpack")
//	// Middleware
//	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
//		Format: "[ECHO] DEBUG method=${method}, uri=${uri}, status=${status}\n",
//	}))
//	e.Use(middleware.Recover())
//	//e.Use(angeldm.NewPOPMiddleware())
//	e.Use(angeldm.NewXormMiddleware())
//	//Set Renderer
//	e.Renderer = angeldm.Default()
//
//	// Routes
//	e.GET("/", func(c echo.Context) error {
//		// POP
//		//cc := c.(*context.CustomContext)
//		//users := models.Users{}
//		//err := cc.Connection.All(&users)
//		//if err != nil {
//		//	panic(err)
//		//}
//		//render 	with master
//		return c.Render(http.StatusOK, "index", echo.Map{
//			"title": "Index title!",
//			"add": func(a int, b int) int {
//				return a + b
//			},
//		})
//	})
//
//	return &Application{Echo: e}
//
//}
//
//func (a *Application) Start() {
//	// Start server
//	a.Echo.Logger.Fatal(a.Echo.Start(":9090"))
//}

package main

import (
	"angeldm.echoview/application"
	"angeldm.echoview/application/ctx"
	"flag"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

var (
	configFlag = flag.String("config",
		"echo-xorm-config.toml",
		"-config=\"path-to-your-config-file\" ")
)

func main() {
	// parse flags
	flag.Parse()

	var (
		err error
		a   *application.Application
	)

	flags := &ctx.Flags{
		CfgFileName: *configFlag,
	}

	// create application
	a, err = application.New(flags)
	if err != nil {
		log.Fatal("error ", os.Args[0]+" initialization error: "+err.Error())
		os.Exit(1)
	}

	// log initialization
	a.Ctx.Logger.Info("ECHO[APP] Application initialized successfully")
	a.Ctx.Logger.WithFields(logrus.Fields{
		"version": a.Ctx.Config.Version,
		"port":    a.Ctx.Config.Port,
		"db":      a.Ctx.Config.Database.Db,
		"dsn":     a.Ctx.Config.Database.Dsn,
	}).Info("ECHO[CONFIGURATION]")

	//.Info("appcontrol", "CONFIG: ",a.Ctx.Config)
	a.Ctx.Logger.WithFields(logrus.Fields{
		"config": flags.CfgFileName,
	}).Info("ECHO[FLAGS]")

	go func() {
		// here we go
		a.Run()
	}()

	// signal control
	sigstop := make(chan os.Signal, 1)
	signal.Notify(sigstop, syscall.SIGTERM, os.Interrupt)

	sig := <-sigstop

	if a.Ctx.Logger != nil {
		a.Ctx.Logger.WithFields(logrus.Fields{
			"exec":   os.Args[0],
			"signal": sig.String(),
		}).Info("ECHO[APP] Caught")
	}

	// shutdown server on signal
	err = a.Shutdown()
	if err != nil {
		if a.Ctx.Logger != nil {
			a.Ctx.Logger.Error("error stopping server", err)
		}
		os.Exit(1)
	}

}

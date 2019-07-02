package logmatic

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var printf = fmt.Printf
var exit = os.Exit

// logFunc represents a log function
type logFunc func(a ...interface{}) string

// Logger maintains a set of logging functions
// and has a log level that can be modified dynamically
type Logger struct {
	ID          string
	level       LogLevel
	trace       logFunc
	debug       logFunc
	info        logFunc
	warn        logFunc
	sql         logFunc
	error       logFunc
	fatal       logFunc
	ExitOnFatal bool
}

type LogLevel uint8

// Log levels
const (
	TRACE = iota
	SQL
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func (l *Logger) now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (l *Logger) Log(level string, format string, a ...interface{}) {
	printf("[%s] %s %s %s\n",
		color.YellowString(l.ID),
		level,
		color.MagentaString("=>"),
		fmt.Sprintf(format, a...))
}

// SetLevel updates the logging level for future logs
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Trace logs a trace statement
// TRACE only
func (l *Logger) Trace(format string, a ...interface{}) {
	if l.level == TRACE {
		l.Log(l.trace("TRACE"), format, a...)
	}
}

// Debug logs a debug statement
// DEBUG or lower
func (l *Logger) Debug(format string, a ...interface{}) {
	if l.level <= DEBUG {
		l.Log(l.debug("DEBUG"), format, a...)
	}
}

// Debug logs a debug statement
// DEBUG or lower
func (l *Logger) Sql(format string, a ...interface{}) {
	if l.level <= SQL {
		l.Log(l.sql("SQL"), format, a...)
	}
}

// Info logs an info statement
// INFO or lower
func (l *Logger) Info(format string, a ...interface{}) {
	if l.level <= INFO {
		l.Log(l.info("INFO"), format, a...)
	}
}

// Warn logs a warn statement
// WARN or lower
func (l *Logger) Warn(format string, a ...interface{}) {
	if l.level <= WARN {
		l.Log(l.warn("WARN"), format, a...)
	}
}

// Error logs an error statement
// ERROR or lower (any level)
func (l *Logger) Error(format string, a ...interface{}) {
	if l.level <= ERROR {
		l.Log(l.error("ERROR"), format, a...)
	}
}

// Fatal logs an error statement and exits the application
// FATAL or lower (any level)
func (l *Logger) Fatal(format string, a ...interface{}) {
	l.Log(l.fatal("FATAL"), format, a...)

	if l.ExitOnFatal {
		exit(1)
	}
}

// NewLogger creates a new logger
// Default level is INFO
func NewLogger(id string) *Logger {
	return &Logger{
		ID:          id,
		level:       INFO,
		trace:       color.New(color.FgBlue).SprintFunc(),
		debug:       color.New(color.FgGreen).SprintFunc(),
		info:        color.New(color.FgCyan).SprintFunc(),
		sql:         color.New(color.FgGreen).SprintFunc(),
		warn:        color.New(color.FgYellow).SprintFunc(),
		error:       color.New(color.FgRed).SprintFunc(),
		fatal:       color.New(color.FgRed, color.Bold).SprintFunc(),
		ExitOnFatal: true,
	}
}

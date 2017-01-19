// Package log implements a common logging like log4j.
package log

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

// std is a default logger for console
var std = New(os.Stdout).SkipCaller(3)

func init() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		std.SetFlags(DEFAULT_FLAGS | F_COLOR)
	}
}

// Get log level
func GetLevel() int {
	return std.GetLevel()
}

// Set log level
func SetLevel(level int) {
	std.SetLevel(level)
}

// Get log level name
func GetLevelName() string {
	return std.GetLevelName()
}

// Set log level name
func SetLevelName(level string) {
	std.SetLevelName(level)
}

// Get the flags for output format
func GetFlags() int {
	return std.GetFlags()
}

// Set flags for output format
func SetFlags(flags int) {
	std.SetFlags(flags)
}

// Get the process name
func GetAppName() string {
	return std.GetAppName()
}

// Set a process name
func SetAppName(name string) {
	std.SetAppName(name)
}

// Get time format for log line
func GetTimeFormat() string {
	return std.GetTimeFormat()
}

// Set time format for log line
func SetTimeFormat(format string) {
	std.SetTimeFormat(format)
}

// Set a writer
func SetWriter(w io.Writer) {
	std.SetWriter(w)
}

// Indicate whether output debug message
func IsDebugEnabled() bool {
	return std.IsDebugEnabled()
}

// Indicate whether output info message
func IsInfoEnabled() bool {
	return std.IsInfoEnabled()
}

// Indicate whether output warning message
func IsWarnEnabled() bool {
	return std.IsWarnEnabled()
}

// Indicate whether output error message
func IsErrorEnabled() bool {
	return std.IsErrorEnabled()
}

// Output a debug message
func Debug(obj ...interface{}) {
	std.Debug(obj...)
}

// Output an info message
func Info(obj ...interface{}) {
	std.Info(obj...)
}

// Output a warning message
func Warn(obj ...interface{}) {
	std.Warn(obj...)
}

// Output an error message
func Error(obj ...interface{}) {
	std.Error(obj...)
}

// Output a fatal message with full stack
func Fatal(obj ...interface{}) {
	std.Fatal(obj...)
}

// Output a debug message
func Debugf(msg string, args ...interface{}) {
	std.Debugf(msg, args...)
}

// Output an info message
func Infof(msg string, args ...interface{}) {
	std.Infof(msg, args...)
}

// Output a warning message
func Warnf(msg string, args ...interface{}) {
	std.Warnf(msg, args...)
}

// Output an error message
func Errorf(msg string, args ...interface{}) {
	std.Errorf(msg, args...)
}

// Output a fatal message with full stack
func Fatalf(msg string, args ...interface{}) {
	std.Fatalf(msg, args...)
}

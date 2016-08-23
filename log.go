package log

import (
	"io"
	"os"
)

// std is a default logger for console
var std = New(os.Stdout).EnableColorizedLevel(true).SkipCaller(3)

// Get log level
func GetLevel() int {
	return std.GetLevel()
}

// Set log level
func SetLevel(level int) {
	std.SetLevel(level)
}

// Set a name to indicate a process
func SetName(name string) {
	std.SetName(name)
}

// Set time layout for log line
func SetTimeLayout(layout string) {
	std.SetTimeLayout(layout)
}

// EnableGoroutineId output goroutinue id
func EnableGoroutineId(enable bool) {
	std.EnableGoroutineId(enable)
}

// EnableLongFileFormat output long file name
func EnableLongFileFormat(enable bool) {
	std.EnableLongFileFormat(enable)
}

// EnableColorizedLevel output colorized level
func EnableColorizedLevel(enable bool) {
	std.EnableColorizedLevel(enable)
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

// Output an debug message
func Debug(msg string, args ...interface{}) {
	std.Debug(msg, args...)
}

// Output an info message
func Info(msg string, args ...interface{}) {
	std.Info(msg, args...)
}

// Output a warning message
func Warn(msg string, args ...interface{}) {
	std.Warn(msg, args...)
}

// Output an error message
func Error(msg string, args ...interface{}) {
	std.Error(msg, args)
}

// Output a fatal message with full stack
func Fatal(msg string, args ...interface{}) {
	std.Fatal(msg, args)
}

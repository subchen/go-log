package log

// Default is a default Logger instance
var Default = New()

// Indicate whether output debug message
func IsDebugEnabled() bool {
	return Default.IsDebugEnabled()
}

// Indicate whether output info message
func IsInfoEnabled() bool {
	return Default.IsInfoEnabled()
}

// Indicate whether output warning message
func IsWarnEnabled() bool {
	return Default.IsWarnEnabled()
}

// Indicate whether output error message
func IsErrorEnabled() bool {
	return Default.IsErrorEnabled()
}

// Indicate whether output fatal message
func IsFatalEnabled() bool {
	return Default.IsFatalEnabled()
}

// Indicate whether output is off
func IsDisabled() bool {
	return Default.IsDisabled()
}

// Output a debug message
func Debug(obj ...interface{}) {
	Default.Debug(obj...)
}

// Output an info message
func Info(obj ...interface{}) {
	Default.Info(obj...)
}

// Output an info message
func Print(obj ...interface{}) {
	Default.Print(obj...)
}

// Output a warning message
func Warn(obj ...interface{}) {
	Default.Warn(obj...)
}

// Output an error message
func Error(obj ...interface{}) {
	Default.Error(obj...)
}

// Output a panic message with full stack
func Panic(obj ...interface{}) {
	Default.Panic(obj...)
}

// Output a fatal message with full stack
func Fatal(obj ...interface{}) {
	Default.Fatal(obj...)
}

// Output a debug message
func Debugln(obj ...interface{}) {
	Default.Debugln(obj...)
}

// Output an info message
func Infoln(obj ...interface{}) {
	Default.Infoln(obj...)
}

// Output an info message
func Println(obj ...interface{}) {
	Default.Println(obj...)
}

// Output a warning message
func Warnln(obj ...interface{}) {
	Default.Warnln(obj...)
}

// Output an error message
func Errorln(obj ...interface{}) {
	Default.Errorln(obj...)
}

// Output a panic message with full stack
func Panicln(obj ...interface{}) {
	Default.Panicln(obj...)
}

// Output a fatal message with full stack
func Fatalln(obj ...interface{}) {
	Default.Fatalln(obj...)
}

// Output a debug message
func Debugf(msg string, args ...interface{}) {
	Default.Debugf(msg, args...)
}

// Output an info message
func Infof(msg string, args ...interface{}) {
	Default.Infof(msg, args...)
}

// Output an info message
func Printf(msg string, args ...interface{}) {
	Default.Printf(msg, args...)
}

// Output a warning message
func Warnf(msg string, args ...interface{}) {
	Default.Warnf(msg, args...)
}

// Output an error message
func Errorf(msg string, args ...interface{}) {
	Default.Errorf(msg, args...)
}

// Output a panic message with full stack
func Panicf(msg string, args ...interface{}) {
	Default.Panicf(msg, args...)
}

// Output a fatal message with full stack
func Fatalf(msg string, args ...interface{}) {
	Default.Fatalf(msg, args...)
}

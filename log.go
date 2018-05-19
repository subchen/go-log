package log

// Default is a default Logger instance
var Default = New()

// IsDebugEnabled indicates whether output message
func IsDebugEnabled() bool {
	return Default.IsDebugEnabled()
}

// IsInfoEnabled indicates whether output message
func IsInfoEnabled() bool {
	return Default.IsInfoEnabled()
}

// IsPrintEnabled indicates whether output message
func IsPrintEnabled() bool {
	return Default.IsPrintEnabled()
}

// IsWarnEnabled indicates whether output message
func IsWarnEnabled() bool {
	return Default.IsWarnEnabled()
}

// IsErrorEnabled indicates whether output message
func IsErrorEnabled() bool {
	return Default.IsErrorEnabled()
}

// IsPanicEnabled indicates whether output message
func IsPanicEnabled() bool {
	return Default.IsPanicEnabled()
}

// IsFatalEnabled indicates whether output message
func IsFatalEnabled() bool {
	return Default.IsFatalEnabled()
}

// IsDisabled indicates whether output message
func IsDisabled() bool {
	return Default.IsDisabled()
}

// Debug outputs message, Arguments are handled by fmt.Sprint
func Debug(obj ...interface{}) {
	Default.Debug(obj...)
}

// Info outputs message, Arguments are handled by fmt.Sprint
func Info(obj ...interface{}) {
	Default.Info(obj...)
}

// Print outputs message, Arguments are handled by fmt.Sprint
func Print(obj ...interface{}) {
	Default.Print(obj...)
}

// Warn outputs message, Arguments are handled by fmt.Sprint
func Warn(obj ...interface{}) {
	Default.Warn(obj...)
}

// Error outputs message, Arguments are handled by fmt.Sprint
func Error(obj ...interface{}) {
	Default.Error(obj...)
}

// Panic outputs message, and followed by a call to panic() Arguments are handled by fmt.Sprint
func Panic(obj ...interface{}) {
	Default.Panic(obj...)
}

// Fatal outputs message, and followed by a call to os.Exit(1) Arguments are handled by fmt.Sprint
func Fatal(obj ...interface{}) {
	Default.Fatal(obj...)
}

// Debugln outputs message, Arguments are handled by fmt.Sprintln
func Debugln(obj ...interface{}) {
	Default.Debugln(obj...)
}

// Infoln outputs message, Arguments are handled by fmt.Sprintln
func Infoln(obj ...interface{}) {
	Default.Infoln(obj...)
}

// Println outputs message, Arguments are handled by fmt.Sprintln
func Println(obj ...interface{}) {
	Default.Println(obj...)
}

// Warnln outputs message, Arguments are handled by fmt.Sprintln
func Warnln(obj ...interface{}) {
	Default.Warnln(obj...)
}

// Errorln outputs message, Arguments are handled by fmt.Sprintln
func Errorln(obj ...interface{}) {
	Default.Errorln(obj...)
}

// Panicln outputs message and followed by a call to panic(), Arguments are handled by fmt.Sprintln
func Panicln(obj ...interface{}) {
	Default.Panicln(obj...)
}

// Fatalln outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprintln
func Fatalln(obj ...interface{}) {
	Default.Fatalln(obj...)
}

// Debugf outputs message, Arguments are handled by fmt.Sprintf
func Debugf(msg string, args ...interface{}) {
	Default.Debugf(msg, args...)
}

// Infof outputs message, Arguments are handled by fmt.Sprintf
func Infof(msg string, args ...interface{}) {
	Default.Infof(msg, args...)
}

// Printf outputs message, Arguments are handled by fmt.Sprintf
func Printf(msg string, args ...interface{}) {
	Default.Printf(msg, args...)
}

// Warnf outputs message, Arguments are handled by fmt.Sprintf
func Warnf(msg string, args ...interface{}) {
	Default.Warnf(msg, args...)
}

// Errorf outputs message, Arguments are handled by fmt.Sprintf
func Errorf(msg string, args ...interface{}) {
	Default.Errorf(msg, args...)
}

// Panicf outputs message and followed by a call to panic(), Arguments are handled by fmt.Sprintf
func Panicf(msg string, args ...interface{}) {
	Default.Panicf(msg, args...)
}

// Fatalf outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprintf
func Fatalf(msg string, args ...interface{}) {
	Default.Fatalf(msg, args...)
}

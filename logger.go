package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Exit is equals os.Exit
var Exit = os.Exit

// Logger is represents an active logging object
type Logger struct {
	m         sync.Mutex
	Level     Level
	Formatter Formatter
	Out       io.Writer
}

// New creates a new Logger
func New() *Logger {
	return &Logger{
		Level:     INFO,
		Formatter: new(simpleFormatter),
		Out:       os.Stdout,
	}
}

// IsDebugEnabled indicates whether output message
func (l *Logger) IsDebugEnabled() bool {
	return l.Level >= DEBUG
}

// IsInfoEnabled indicates whether output message
func (l *Logger) IsInfoEnabled() bool {
	return l.Level >= INFO
}

// IsPrintEnabled indicates whether output message
func (l *Logger) IsPrintEnabled() bool {
	return l.Level > OFF
}

// IsWarnEnabled indicates whether output message
func (l *Logger) IsWarnEnabled() bool {
	return l.Level >= WARN
}

// IsErrorEnabled indicates whether output message
func (l *Logger) IsErrorEnabled() bool {
	return l.Level >= ERROR
}

// IsPanicEnabled indicates whether output message
func (l *Logger) IsPanicEnabled() bool {
	return l.Level >= PANIC
}

// IsFatalEnabled indicates whether output message
func (l *Logger) IsFatalEnabled() bool {
	return l.Level >= FATAL
}

// IsDisabled indicates whether output message
func (l *Logger) IsDisabled() bool {
	return l.Level <= OFF
}

// Debug outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Debug(obj ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprint(obj...))
	}
}

// Info outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Info(obj ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprint(obj...))
	}
}

// Print outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Print(obj ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, fmt.Sprint(obj...))
	}
}

// Warn outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Warn(obj ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprint(obj...))
	}
}

// Error outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Error(obj ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprint(obj...))
	}
}

// Panic outputs message, and followed by a call to panic() Arguments are handled by fmt.Sprint
func (l *Logger) Panic(obj ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprint(obj...))
	}
	panic(fmt.Sprint(obj...))
}

// Fatal outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprint
func (l *Logger) Fatal(obj ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, fmt.Sprint(obj...))
	}
	Exit(1)
}

// Debugln outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Debugln(obj ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, vsprintln(obj...))
	}
}

// Infoln outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Infoln(obj ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, vsprintln(obj...))
	}
}

// Println outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Println(obj ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, vsprintln(obj...))
	}
}

// Warnln outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Warnln(obj ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, vsprintln(obj...))
	}
}

// Errorln outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Errorln(obj ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, vsprintln(obj...))
	}
}

// Panicln outputs message and followed by a call to panic(), Arguments are handled by fmt.Sprintln
func (l *Logger) Panicln(obj ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, vsprintln(obj...))
	}
	panic(vsprintln(obj...))
}

// Fatalln outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprintln
func (l *Logger) Fatalln(obj ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, vsprintln(obj...))
	}
	Exit(1)
}

// Debugf outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Debugf(msg string, args ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprintf(msg, args...))
	}
}

// Infof outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Infof(msg string, args ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprintf(msg, args...))
	}
}

// Printf outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Printf(msg string, args ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, fmt.Sprintf(msg, args...))
	}
}

// Warnf outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Warnf(msg string, args ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprintf(msg, args...))
	}
}

// Errorf outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Errorf(msg string, args ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprintf(msg, args...))
	}
}

// Panicf outputs message and followed by a call to panic(), Arguments are handles by fmt.Sprintf
func (l *Logger) Panicf(msg string, args ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprintf(msg, args...))
	}
	panic(fmt.Sprintf(msg, args...))
}

// Fatalf outputs message and followed by a call to os.Exit(1), Arguments are handles by fmt.Sprintf
func (l *Logger) Fatalf(msg string, args ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, fmt.Sprintf(msg, args...))
	}
	Exit(1)
}

func (l *Logger) log(level Level, msg string) {
	line := l.Formatter.Format(level, msg, l)

	l.m.Lock()
	defer l.m.Unlock()

	_, err := l.Out.Write(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write log, %v\n", err)
	}
}

// vsprintln => spaces are always added between operands
func vsprintln(obj ...interface{}) string {
	msg := fmt.Sprintln(obj...)
	return msg[:len(msg)-1]
}

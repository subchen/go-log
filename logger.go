package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var Exit = os.Exit

type Logger struct {
	m         sync.Mutex
	Level     Level
	Formatter Formatter
	Out       io.Writer
}

func New() *Logger {
	return &Logger{
		Level:     INFO,
		Formatter: new(simpleFormatter),
		Out:       os.Stdout,
	}
}

func (l *Logger) IsDebugEnabled() bool {
	return l.Level >= DEBUG
}

func (l *Logger) IsInfoEnabled() bool {
	return l.Level >= INFO
}

func (l *Logger) IsWarnEnabled() bool {
	return l.Level >= WARN
}

func (l *Logger) IsErrorEnabled() bool {
	return l.Level >= ERROR
}

func (l *Logger) IsFatalEnabled() bool {
	return l.Level >= FATAL
}

func (l *Logger) IsDisabled() bool {
	return l.Level >= OFF
}

func (l *Logger) Debug(obj ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprint(obj...))
	}
}

func (l *Logger) Info(obj ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprint(obj...))
	}
}

func (l *Logger) Print(obj ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, fmt.Sprint(obj...))
	}
}

func (l *Logger) Warn(obj ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprint(obj...))
	}
}

func (l *Logger) Error(obj ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprint(obj...))
	}
}

func (l *Logger) Panic(obj ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprint(obj...))
	}
	panic(fmt.Sprint(obj...))
}

func (l *Logger) Fatal(obj ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, fmt.Sprint(obj...))
	}
	Exit(1)
}

func (l *Logger) Debugln(obj ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, vsprintln(obj...))
	}
}

func (l *Logger) Infoln(obj ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, vsprintln(obj...))
	}
}

func (l *Logger) Println(obj ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, vsprintln(obj...))
	}
}

func (l *Logger) Warnln(obj ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, vsprintln(obj...))
	}
}

func (l *Logger) Errorln(obj ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, vsprintln(obj...))
	}
}

func (l *Logger) Panicln(obj ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, vsprintln(obj...))
	}
	panic(vsprintln(obj...))
}

func (l *Logger) Fatalln(obj ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, vsprintln(obj...))
	}
	Exit(1)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Printf(msg string, args ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Panicf(msg string, args ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprintf(msg, args...))
	}
	panic(fmt.Sprintf(msg, args...))
}

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

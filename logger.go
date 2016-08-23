package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/subchen/gstack/gls"
)

// Log Level
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

var (
	levelStr = []string{
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL",
	}
	levelStrWithColor = []string{
		"\033[34mDEBUG\033[0m",
		"\033[32mINFO\033[0m",
		"\033[33mWARN\033[0m",
		"\033[31mERROR\033[0m",
		"\033[35mFATAL\033[0m",
	}
)

func New(out io.Writer) *Logger {
	return &Logger{
		out:            out,
		level:          INFO,
		pid:            os.Getpid(),
		name:           "",
		timeLayout:     "2006-01-02 15:04:05.000",
		goroutineId:    false,
		longFileFormat: false,
		colorizedLevel: false,
		callerSkip:     2,
	}
}

type Logger struct {
	mu             sync.Mutex
	out            io.Writer
	level          int
	pid            int
	name           string
	timeLayout     string
	goroutineId    bool
	longFileFormat bool
	colorizedLevel bool
	callerSkip     int
}

func (l *Logger) GetLevel() int {
	return l.level
}

func (l *Logger) SetLevel(level int) *Logger {
	l.level = level
	return l
}

func (l *Logger) SetName(name string) *Logger {
	l.name = name
	return l
}

func (l *Logger) SetTimeLayout(layout string) *Logger {
	l.timeLayout = layout
	return l
}

func (l *Logger) EnableGoroutineId(enable bool) *Logger {
	l.goroutineId = enable
	return l
}

func (l *Logger) EnableLongFileFormat(enable bool) *Logger {
	l.longFileFormat = enable
	return l
}

func (l *Logger) EnableColorizedLevel(enable bool) *Logger {
	l.colorizedLevel = enable
	return l
}

func (l *Logger) SkipCaller(skip int) *Logger {
	l.callerSkip = skip
	return l
}

func (l *Logger) SetWriter(w io.Writer) *Logger {
	l.out = w
	return l
}

func (l *Logger) IsDebugEnabled() bool {
	return l.level <= DEBUG
}

func (l *Logger) IsInfoEnabled() bool {
	return l.level <= INFO
}

func (l *Logger) IsWarnEnabled() bool {
	return l.level <= WARN
}

func (l *Logger) IsErrorEnabled() bool {
	return l.level <= ERROR
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.level <= DEBUG {
		l.log(DEBUG, msg, args...)
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l.level <= INFO {
		l.log(INFO, msg, args...)
	}
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	if l.level <= WARN {
		l.log(WARN, msg, args...)
	}
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l.level <= ERROR {
		l.log(ERROR, msg, args...)
	}
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.log(FATAL, msg, args...)
}

func (l *Logger) log(level int, msg string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(l.callerSkip)
	if !ok {
		file = "???"
		line = 0
	} else if !l.longFileFormat {
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	}

	// output format: DATE PID [NAME] [GID] LEVEL file:line message
	// 2001-10-10 12:00:00,000+0800 1234 app 987 INFO main.go:1234 log message ...
	buf := new(bytes.Buffer)
	buf.WriteByte(' ')
	buf.WriteString(strconv.Itoa(l.pid))
	buf.WriteByte(' ')
	if l.name != "" {
		buf.WriteString(l.name)
		buf.WriteByte(' ')
	}
	if l.goroutineId || l.level == DEBUG {
		buf.WriteString(strconv.FormatUint(gls.GoroutineID(), 10))
		buf.WriteByte(' ')
	}
	if l.colorizedLevel {
		buf.WriteString(levelStrWithColor[level])
	} else {
		buf.WriteString(levelStr[level])
	}
	buf.WriteByte(' ')
	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	buf.WriteByte(' ')
	fmt.Fprintf(buf, msg, args...)
	buf.WriteByte('\n')

	if level == FATAL {
		for i := l.callerSkip; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			fmt.Fprintf(buf, "\tat %s:%d (0x%x)\n", file, line, pc)
		}
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	timeStr := time.Now().Format(l.timeLayout)

	l.out.Write([]byte(timeStr))
	l.out.Write(buf.Bytes())
}

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

	"github.com/subchen/gls"
)

const (
	// log level
	L_DEBUG = iota
	L_INFO
	L_WARN
	L_ERROR
	L_FATAL
	L_OFF

	// Bits or'ed together to control what's printed.
	F_TIME = 1 << iota
	F_LONG_FILE
	F_SHORT_FILE
	F_PID
	F_GID
	F_COLOR

	// default flags
	DEFAULT_FLAGS = F_TIME | F_SHORT_FILE | F_PID

	// default time format
	DEFAULT_TIME_FORMAT = "2006-01-02 15:04:05.000"
)

var (
	levelStr = []string{
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL",
		"OFF",
	}
	levelStrWithColor = []string{
		"\033[34mDEBUG\033[0m",
		"\033[32mINFO\033[0m",
		"\033[33mWARN\033[0m",
		"\033[31mERROR\033[0m",
		"\033[35mFATAL\033[0m",
		"OFF",
	}

	buffer = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

func New(out io.Writer) *Logger {
	return &Logger{
		writer:     out,
		level:      L_INFO,
		pid:        os.Getpid(),
		name:       "",
		timeFormat: DEFAULT_TIME_FORMAT,
		flags:      DEFAULT_FLAGS,
		callerSkip: 2,
	}
}

type Logger struct {
	m          sync.Mutex
	writer     io.Writer
	level      int
	pid        int
	name       string
	timeFormat string
	flags      int
	callerSkip int
}

func (l *Logger) GetLevel() int {
	return l.level
}

func (l *Logger) SetLevel(level int) *Logger {
	if level < L_DEBUG || level > L_FATAL {
		panic("unknown log level")
	}

	l.m.Lock()
	l.level = level
	l.m.Unlock()
	return l
}

func (l *Logger) GetLevelName() string {
	return levelStr[l.level]
}

func (l *Logger) SetLevelName(level string) {
	level = strings.ToUpper(level)
	for i, v := range levelStr {
		if v == level {
			l.SetLevel(i)
			return
		}
	}

	panic("unknown log level: " + level)
}

func (l *Logger) GetFlags() int {
	return l.flags
}

func (l *Logger) SetFlags(flags int) *Logger {
	l.m.Lock()
	l.flags = flags
	l.m.Unlock()
	return l
}

func (l *Logger) GetAppName() string {
	return l.name
}

func (l *Logger) SetAppName(name string) *Logger {
	l.m.Lock()
	l.name = name
	l.m.Unlock()
	return l
}

func (l *Logger) GetTimeFormat() string {
	return l.timeFormat
}

func (l *Logger) SetTimeFormat(format string) *Logger {
	l.m.Lock()
	l.timeFormat = format
	l.m.Unlock()
	return l
}

func (l *Logger) SkipCaller(skip int) *Logger {
	l.m.Lock()
	defer l.m.Unlock()
	l.callerSkip = skip
	return l
}

func (l *Logger) SetWriter(w io.Writer) *Logger {
	l.m.Lock()
	defer l.m.Unlock()
	l.writer = w
	return l
}

func (l *Logger) IsDebugEnabled() bool {
	return l.level <= L_DEBUG
}

func (l *Logger) IsInfoEnabled() bool {
	return l.level <= L_INFO
}

func (l *Logger) IsWarnEnabled() bool {
	return l.level <= L_WARN
}

func (l *Logger) IsErrorEnabled() bool {
	return l.level <= L_ERROR
}

func (l *Logger) IsFatalEnabled() bool {
	return l.level <= L_FATAL
}

func (l *Logger) Debug(obj ...interface{}) {
	if l.level <= L_DEBUG {
		l.log(L_DEBUG, fmt.Sprint(obj...))
	}
}

func (l *Logger) Info(obj ...interface{}) {
	if l.level <= L_INFO {
		l.log(L_INFO, fmt.Sprint(obj...))
	}
}

func (l *Logger) Warn(obj ...interface{}) {
	if l.level <= L_WARN {
		l.log(L_WARN, fmt.Sprint(obj...))
	}
}

func (l *Logger) Error(obj ...interface{}) {
	if l.level <= L_ERROR {
		l.log(L_ERROR, fmt.Sprint(obj...))
	}
}

func (l *Logger) Fatal(obj ...interface{}) {
	if l.level <= L_FATAL {
		l.log(L_FATAL, fmt.Sprint(obj...))
	}
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	if l.level <= L_DEBUG {
		l.log(L_DEBUG, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	if l.level <= L_INFO {
		l.log(L_INFO, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	if l.level <= L_WARN {
		l.log(L_WARN, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	if l.level <= L_ERROR {
		l.log(L_ERROR, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	if l.level <= L_FATAL {
		l.log(L_FATAL, fmt.Sprintf(msg, args...))
	}
}

func (l *Logger) log(level int, msg string) {
	// output format: DATE PID [NAME] [GID] LEVEL file:line message
	// 2001-10-10 12:00:00,000+0800 1234 app 987 INFO main.go:1234 log message ...

	buf := buffer.Get().(*bytes.Buffer)
	defer buffer.Put(buf)

	if l.flags&F_TIME != 0 {
		timeStr := time.Now().Format(l.timeFormat)
		buf.WriteString(timeStr)
		buf.WriteByte(' ')
	}
	if l.flags&F_PID != 0 {
		buf.WriteString(strconv.Itoa(l.pid))
		buf.WriteByte(' ')
	}
	if l.name != "" {
		buf.WriteString(l.name)
		buf.WriteByte(' ')
	}
	if l.flags&F_GID != 0 {
		buf.WriteString(strconv.FormatUint(gls.GoroutineID(), 10))
		buf.WriteByte(' ')
	}
	if l.flags&F_COLOR != 0 {
		buf.WriteString(levelStrWithColor[level])
	} else {
		buf.WriteString(levelStr[level])
	}
	buf.WriteByte(' ')

	if l.flags&(F_LONG_FILE|F_SHORT_FILE) != 0 {
		_, file, line, ok := runtime.Caller(l.callerSkip)
		if !ok {
			file = "???"
			line = 0
		} else if l.flags&F_SHORT_FILE != 0 {
			if index := strings.LastIndex(file, "/"); index >= 0 {
				file = file[index+1:]
			} else if index = strings.LastIndex(file, "\\"); index >= 0 {
				file = file[index+1:]
			}
		}
		buf.WriteString(file)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
		buf.WriteByte(' ')
	}
	buf.WriteString(msg)
	buf.WriteByte('\n')

	if level == L_FATAL {
		for i := l.callerSkip; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			fmt.Fprintf(buf, "\tat %s:%d (0x%x)\n", file, line, pc)
		}
	}

	line := buf.Bytes()
	buf.Reset()

	l.m.Lock()
	defer l.m.Unlock()

	l.writer.Write(line)
}

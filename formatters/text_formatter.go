package formatters

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/subchen/go-log"
)

var (
	fmtBuffer = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

// TextFormatter is a text line formatter
type TextFormatter struct {
	AppName    string
	TimeFormat string

	init   sync.Once
	host   []byte
	app    []byte
	pid    []byte
	isterm bool
}

// Format implements log.Formatter
func (f TextFormatter) Format(level log.Level, msg string, logger *log.Logger) []byte {
	// output format: DATE LEVEL HOST APP PID file:line message
	// 2001-10-10T12:00:00,000+0800 INFO web-1 app 1234 main/main.go:1234 message ...

	f.init.Do(func() {
		if f.AppName == "" {
			f.AppName = filepath.Base(os.Args[0])
		}
		f.app = []byte(f.AppName)

		if f.TimeFormat == "" {
			f.TimeFormat = "2006-01-02T15:04:05.000-0700"
		}

		f.isterm = IsTerminal(logger.Out)

		host, _ = os.Hostname()
		f.host = []byte(host)

		f.pid = []byte(strconv.Itoa(os.Getpid()))
	})

	buf := fmtBuffer.Get().(*bytes.Buffer)
	buf.Reset()
	defer fmtBuffer.Put(buf)

	// timestamp
	timeStr := time.Now().Format(f.TimeFormat)
	buf.WriteString(timeStr)

	// level
	buf.WriteByte(' ')
	if f.isterm {
		buf.WriteString(level.ColorString())
	} else {
		buf.WriteString(level.String())
	}

	// host
	buf.WriteByte(' ')
	buf.Write(f.host)

	// name
	buf.WriteByte(' ')
	buf.Write(f.app)

	// pid
	buf.WriteByte(' ')
	buf.Write(f.pid)

	// file, line
	file, line := FilelineCaller(7)
	buf.WriteByte(' ')
	buf.WriteString(file)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	buf.WriteByte('"')

	// msg
	buf.WriteString(msg)

	// newline
	buf.WriteByte('\n')

	return buf.Bytes()
}

package formatters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/subchen/go-log"
)

// JSONFormatter is a json formatter
type JSONFormatter struct {
	AppName    string
	TimeFormat string

	init sync.Once
	host string
	pid  int
}

// Format implements log.Formatter
func (f JSONFormatter) Format(level log.Level, msg string, logger *log.Logger) []byte {
	// output fields: time level host app pid file line msg

	f.init.Do(func() {
		if f.AppName == "" {
			f.AppName = filepath.Base(os.Args[0])
		}
		if f.TimeFormat == "" {
			f.TimeFormat = "2006-01-02T15:04:05.000-0700"
		}

		f.host, _ = os.Hostname()
		f.pid = os.Getpid()
	})

	data := make(map[string]interface{}, 8)

	// file, line
	file, line := FilelineCaller(7)

	data["time"] = time.Now().Format(f.TimeFormat)
	data["level"] = level.String()
	data["host"] = f.host
	data["app"] = f.AppName
	data["pid"] = f.pid
	data["file"] = file
	data["line"] = line
	data["msg"] = msg

	line, err := marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal json, %v\n", err)
	}
	return line
}

func marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

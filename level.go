package log

import (
	"fmt"
	"strings"
)

// Level type
type Level uint32

// These are the different logging levels
const (
	OFF Level = iota
	FATAL
	PANIC
	ERROR
	WARN
	INFO
	DEBUG
)

// String converts the Level to a string
func (level Level) String() string {
	switch level {
	case OFF:
		return "OFF"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

// ColorString converts the Level to a string with term colorful
func (level Level) ColorString() string {
	switch level {
	case OFF:
		return "OFF"
	case FATAL:
		return "\033[35mFATAL\033[0m"
	case PANIC:
		return "\033[35mPANIC\033[0m"
	case ERROR:
		return "\033[31mERROR\033[0m"
	case WARN:
		return "\033[33mWARN\033[0m"
	case INFO:
		return "\033[32mINFO\033[0m"
	case DEBUG:
		return "\033[34mDEBUG\033[0m"
	default:
		return "UNKNOWN"
	}
}

// ParseLevel takes a string level and returns the log level constant.
func ParseLevel(name string) (Level, error) {
	switch strings.ToUpper(name) {
	case "OFF":
		return OFF, nil
	case "FATAL":
		return FATAL, nil
	case "PANIC":
		return PANIC, nil
	case "ERROR":
		return ERROR, nil
	case "WARN":
		return WARN, nil
	case "INFO":
		return INFO, nil
	case "DEBUG":
		return DEBUG, nil
	}

	return 0, fmt.Errorf("invalid log.Level: %q", name)
}

package log

import (
	"fmt"
	"time"
)

// Formatter is a interface used to implement a custom Formatter
type Formatter interface {
	Format(level Level, msg string, logger *Logger) []byte
}

// simpleFormatter is default formmatter
type simpleFormatter struct {
}

// Format implements log.Formatter
func (f *simpleFormatter) Format(level Level, msg string, logger *Logger) []byte {
	time := time.Now().Format("15:04:05.000")
	return []byte(fmt.Sprintf("%s %s %s\n", time, level.String(), msg))
}

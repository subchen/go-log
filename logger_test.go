package log

import (
	"bytes"
	"testing"
)

var hits int = 0

type hitsFormatter struct{}

func (f *hitsFormatter) Format(level Level, msg string, logger *Logger) []byte {
	hits++
	return []byte(msg)
}

var l = &Logger{
	Level:     DEBUG,
	Formatter: new(hitsFormatter),
	Out:       new(bytes.Buffer),
}

func reset() {
	l.Level = DEBUG
	hits = 0
}

func TestLogOnLevel(t *testing.T) {
	tests := []struct {
		level Level
		fn    func(...interface{})
		hits  int
	}{
		{level: DEBUG, fn: l.Debug, hits: 1},
		{level: DEBUG, fn: l.Info, hits: 1},
		{level: INFO, fn: l.Debug, hits: 0},
		{level: INFO, fn: l.Info, hits: 1},
		{level: WARN, fn: l.Debugln, hits: 0},
		{level: WARN, fn: l.Infoln, hits: 0},
		{level: WARN, fn: l.Println, hits: 1},
		{level: WARN, fn: l.Errorln, hits: 1},
		{level: ERROR, fn: l.Debug, hits: 0},
		{level: ERROR, fn: l.Print, hits: 1},
		{level: ERROR, fn: l.Error, hits: 1},
		{level: OFF, fn: l.Debugln, hits: 0},
		{level: OFF, fn: l.Println, hits: 0},
		{level: OFF, fn: l.Errorln, hits: 0},
	}

	for i, tt := range tests {
		l.Level = tt.level
		hits = 0
		
		tt.fn("message")
		
		if hits != tt.hits {
			t.Errorf("Case %d, fn hits on level %v, got %v, want %v", i, tt.level, hits, tt.hits)
		}
	}
}

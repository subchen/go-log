package log

import (
	"strings"
	"testing"
)

func TestLevelToString(t *testing.T) {
	tests := []struct {
		level   Level
		wantOut string
	}{
		{level: DEBUG, wantOut: "DEBUG"},
		{level: INFO, wantOut: "INFO"},
		{level: WARN, wantOut: "WARN"},
		{level: ERROR, wantOut: "ERROR"},
		{level: PANIC, wantOut: "PANIC"},
		{level: FATAL, wantOut: "FATAL"},
		{level: OFF, wantOut: "OFF"},
	}

	for _, tt := range tests {
		got := tt.level.String()
		if got != tt.wantOut {
			t.Errorf("%v.String() output = %v, want %v", tt.level, got, tt.wantOut)
		}

		gotColor := tt.level.ColorString()
		if !strings.Contains(gotColor, tt.wantOut) {
			t.Errorf("%v.ColorString() output = %v, want %v", tt.level, gotColor, tt.wantOut)
		}
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name    string
		wantOut Level
		wantErr bool
	}{
		{name: "debug", wantOut: DEBUG},
		{name: "Info", wantOut: INFO},
		{name: "WARN", wantOut: WARN},
		{name: "error", wantOut: ERROR},
		{name: "panic", wantOut: PANIC},
		{name: "FATAL", wantOut: FATAL},
		{name: "Off", wantOut: OFF},
		{name: "xxxx", wantOut: 0, wantErr: true},
	}
	for _, tt := range tests {
		got, err := ParseLevel(tt.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseLevel(%q) error = %v, wantErr %v", tt.name, err, tt.wantErr)
			return
		}
		if got != tt.wantOut {
			t.Errorf("ParseLevel(%q) output = %v, want %v", tt.name, got, tt.wantOut)
		}
	}
}

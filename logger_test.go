package log

import (
	"bytes"
	"testing"
)

var (
	logHits  int = 0
	exitHits int = 0
)

type hitsFormatter struct{}

func (f *hitsFormatter) Format(level Level, msg string, logger *Logger) []byte {
	logHits++
	return []byte(msg)
}

func init() {
	Default.Formatter = new(hitsFormatter)
	Default.Out = new(bytes.Buffer)

	Exit = func(code int) {
		exitHits++
	}
}

func TestLogOnLevel(t *testing.T) {
	type fns struct {
		fn1     func(...interface{})
		fn2     func(...interface{})
		fn3     func(string, ...interface{})
		ispanic bool
		isexit  bool
	}
	debugFns := fns{Debug, Debugln, Debugf, false, false}
	infoFns := fns{Info, Infoln, Infof, false, false}
	printFns := fns{Print, Println, Printf, false, false}
	warnFns := fns{Warn, Warnln, Warnf, false, false}
	errorFns := fns{Error, Errorln, Errorf, false, false}
	panicFns := fns{Panic, Panicln, Panicf, true, false}
	fatalFns := fns{Fatal, Fatalln, Fatalf, false, true}

	tests := []struct {
		level Level
		fns   fns
		hits  int
	}{
		{level: DEBUG, fns: debugFns, hits: 1},
		{level: DEBUG, fns: infoFns, hits: 1},
		{level: DEBUG, fns: printFns, hits: 1},
		{level: DEBUG, fns: warnFns, hits: 1},
		{level: DEBUG, fns: errorFns, hits: 1},
		{level: DEBUG, fns: panicFns, hits: 1},
		{level: DEBUG, fns: fatalFns, hits: 1},

		{level: INFO, fns: debugFns, hits: 0},
		{level: INFO, fns: infoFns, hits: 1},
		{level: INFO, fns: printFns, hits: 1},
		{level: INFO, fns: warnFns, hits: 1},
		{level: INFO, fns: errorFns, hits: 1},
		{level: INFO, fns: panicFns, hits: 1},
		{level: INFO, fns: fatalFns, hits: 1},

		{level: WARN, fns: debugFns, hits: 0},
		{level: WARN, fns: infoFns, hits: 0},
		{level: WARN, fns: printFns, hits: 1},
		{level: WARN, fns: warnFns, hits: 1},
		{level: WARN, fns: errorFns, hits: 1},
		{level: WARN, fns: panicFns, hits: 1},
		{level: WARN, fns: fatalFns, hits: 1},

		{level: ERROR, fns: debugFns, hits: 0},
		{level: ERROR, fns: infoFns, hits: 0},
		{level: ERROR, fns: printFns, hits: 1},
		{level: ERROR, fns: warnFns, hits: 0},
		{level: ERROR, fns: errorFns, hits: 1},
		{level: ERROR, fns: panicFns, hits: 1},
		{level: ERROR, fns: fatalFns, hits: 1},

		{level: PANIC, fns: debugFns, hits: 0},
		{level: PANIC, fns: infoFns, hits: 0},
		{level: PANIC, fns: printFns, hits: 1},
		{level: PANIC, fns: warnFns, hits: 0},
		{level: PANIC, fns: errorFns, hits: 0},
		{level: PANIC, fns: panicFns, hits: 1},
		{level: PANIC, fns: fatalFns, hits: 1},

		{level: FATAL, fns: debugFns, hits: 0},
		{level: FATAL, fns: infoFns, hits: 0},
		{level: FATAL, fns: printFns, hits: 1},
		{level: FATAL, fns: warnFns, hits: 0},
		{level: FATAL, fns: errorFns, hits: 0},
		{level: FATAL, fns: panicFns, hits: 0},
		{level: FATAL, fns: fatalFns, hits: 1},

		{level: OFF, fns: debugFns, hits: 0},
		{level: OFF, fns: infoFns, hits: 0},
		{level: OFF, fns: printFns, hits: 0},
		{level: OFF, fns: warnFns, hits: 0},
		{level: OFF, fns: errorFns, hits: 0},
		{level: OFF, fns: panicFns, hits: 0},
		{level: OFF, fns: fatalFns, hits: 0},
	}

	for i, tt := range tests {
		func() {
			defer func() {
				// check panic hits
				if err := recover(); err != nil {
					if !tt.fns.ispanic {
						t.Errorf("Case %d, got panic", i)
					}
				} else {
					if tt.fns.ispanic {
						t.Errorf("Case %d, no panic found", i)
					}
				}
			}()

			// reset
			logHits = 0
			exitHits = 0
			Default.Level = tt.level

			// run
			tt.fns.fn1("message")
			tt.fns.fn2("message")
			tt.fns.fn3("message")

			// check log hits
			if logHits != tt.hits*3 {
				t.Errorf("Case %d, fn hits on level %v, got %v, want %v", i, tt.level, logHits, tt.hits)
			}

			// check exit hits
			if tt.fns.isexit {
				if exitHits != 3 {
					t.Errorf("Case %d, no exits hits", i)
				}
			}
		}()
	}
}

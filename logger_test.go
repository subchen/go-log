package log

import (
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	stdout := New(os.Stdout)
	stdout.SetTimeLayout("15:04:05.999")
	stdout.SetName("main")
	stdout.SetLevel(DEBUG)
	stdout.EnableColorizedLevel(true)
	stdout.EnableGoroutineId(true)
	for i := 0; i < 3; i++ {
		go func(i int) {
			stdout.Debug("i = %d", i)
			stdout.Info("i = %d", i)
		}(i)
	}
	for i := 0; i < 3; i++ {
		go func(i int) {
			stdout.Debug("i = %d", i)
			stdout.Info("i = %d", i)
		}(i)
	}

	stdout.Warn("warning")
	stdout.Error("error")
	stdout.Fatal("fatal")

	time.Sleep(1 * time.Second)
}

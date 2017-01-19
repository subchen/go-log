package log

import (
	"os"
	"sync"
	"testing"
)

func TestLogger(t *testing.T) {
	stdout := New(os.Stdout)
	stdout.SetTimeFormat("15:04:05.999")
	stdout.SetAppName("main")
	stdout.SetLevel(L_DEBUG)
	stdout.SetFlags(DEFAULT_FLAGS | F_GID | F_COLOR)

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			stdout.Debugf("i = %d", i)
			stdout.Infof("i = %d", i)
			wg.Done()
		}(i)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			stdout.Debugf("i = %d", i)
			stdout.Infof("i = %d", i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	stdout.Warn("warning", "message")
	stdout.Error("error")
	stdout.Fatal("fatal")
}

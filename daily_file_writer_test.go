package log

import (
	"testing"
)

func TestDailyFileWriter(t *testing.T) {
	log := New(&DailyFileWriter{
		Name: "/tmp/test.log",
	})
	log.SetTimeLayout("15:04:05.999")
	log.SetName("main")
	for i := 0; i < 100; i++ {
		log.Info("i = %d", i)
	}
}

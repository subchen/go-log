package log

import (
	"os"
	"testing"
)

func TestCompositeWriters(t *testing.T) {
	log := New(&CompositeWriters{
		os.Stdout,
		&DailyFileWriter{
			Name: "/tmp/test.log",
		},
	})
	log.SetTimeLayout("15:04:05.999")
	log.SetName("main")
	for i := 0; i < 5; i++ {
		log.Info("i = %d", i)
	}
}

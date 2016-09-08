package log

import (
	"testing"
)

func TestFixedSizeFileWriter(t *testing.T) {
	log := New(&FixedSizeFileWriter{
		Name:     "/tmp/test.log",
		MaxSize:  1024 * 3,
		MaxCount: 5,
	})
	log.SetTimeLayout("15:04:05.999")
	log.SetName("main")
	for i := 0; i < 100; i++ {
		log.Info("i = %d", i)
	}
}

package log_test

import (
	"github.com/subchen/go-log"
	"os"
)

func ExampleNew() {
	logger := log.New(os.Stdout)
	logger.SetAppName("app")
	logger.SetFlags(log.DEFAULT_FLAGS | log.F_GID)

	logger.Info("testing ...")
	logger.Errorf("err = %v", os.ErrInvalid)
	// Outputs:
	// 2001-10-10 12:00:00,000 1234 app 987 INFO main.go:13 testing ...
	// 2001-10-10 12:00:00,000 1234 app 987 ERROR main.go:14 err = invalid argument
}

func ExampleAlwaysNewFileWriter() {
	log.SetWriter(&log.AlwaysNewFileWriter{
		Name:     "/tmp/test.log",
		MaxCount: 10,
	})
}

func ExampleDailyFileWriter() {
	log.SetWriter(&log.DailyFileWriter{
		Name:     "/tmp/test.log",
		MaxCount: 10,
	})
}

func ExampleFixedSizeFileWriter() {
	log.SetWriter(&log.FixedSizeFileWriter{
		Name:     "/tmp/test.log",
		MaxSize:  10 * 1024 * 1024, // 10M
		MaxCount: 10,
	})
}

go-log
================

[![Go Report Card](https://goreportcard.com/badge/github.com/subchen/go-log)](https://goreportcard.com/report/github.com/subchen/go-log)
[![GoDoc](https://godoc.org/github.com/subchen/go-log?status.svg)](https://godoc.org/github.com/subchen/go-log)

Logging package similar to log4j for the Golang.

* Support dynamic log level
* Daily log file output
* Fixed size log file ouput
* Output console and file all at once
* Output full stack
* Output goroutinue id to console
* Output colorized level to console

Installation
---------------

```bash
$ go get github.com/subchen/go-log
```

Usage
---------------

### Default log to console

```go
package main

import (
    "os"
    "errors"
    "github.com/subchen/go-log"
)

func main() {
    log.Debugf("app = %s", os.Args[0])
    log.Errorf("error = %v", errors.New("some error"))

    // dynamic set level
    log.SetLevel(log.L_INFO)

    log.Debug("cannot output debug message")
    log.Info("can output info message")
}
```

### Output to file

You can call `SetWriter` to set a file writer into log.

```go
log.SetWriter(&log.FixedSizeFileWriter{
    Name:     "/tmp/test.log",
    MaxSize:  10 * 1024 * 1024, // 10m
    MaxCount: 10,
})
```

### We defined three writers for use

```go
// Create log file if file size large than fixed size (10m)
// files: /tmp/test.log.0 .. test.log.10
&log.FixedSizeFileWriter{
    Name:     "/tmp/test.log",
    MaxSize:  10 * 1024 * 1024, // 10m
    MaxCount: 10,
}

// Create log file every day.
// files: /tmp/test.log.20160102
&log.DailyFileWriter{
    Name: "/tmp/test.log",
    MaxCount: 10,
}

// Create log file every process.
// files: /tmp/test.log.20160102_150405
&log.AlwaysNewFileWriter{
    Name: "/tmp/test.log",
    MaxCount: 10,
}

// Output to multiple writes
io.MultiWriter(
    os.Stdout,
    &log.DailyFileWriter{
        Name: "/tmp/test.log",
        MaxCount: 10,
    }
    //...
)
```

### New log instance

```go
import (
    "github.com/subchen/go-log"
)

func main() {
    logger := log.New(&log.DailyFileWriter{
        Name: "/tmp/test.log",
    })

    logger.Debugf("i = %d", 99)
}
```

### Output stack

You can use `log.Fatal(...)` or `log.Fatalf(...)` to output full stack

```
21:04:32.884 main FATAL logger_test.go:24 this is a fatal message
	at /go/src/github.com/subchen/go-log/logger_test.go:24 (0x81db3)
	at /usr/local/Cellar/go/1.5.2/libexec/src/testing/testing.go:456 (0x786c8)
	at /usr/local/Cellar/go/1.5.2/libexec/src/runtime/asm_amd64.s:1721 (0x59641)
```

### Log format

```
# time pid [name] [gid] level file:line message
2016-02-10 19:33:02.587 12345 main 987 INFO fixed_size_file_writer_test.go:16 message ...
```

* `log.SetTimeFormat("2006-01-02 15:04:05.999")` customize time format
* `log.SetAppName("main")` add a name in log for indicate process
* `log.SetFlags(F_TIME | F_LONG_FILE | F_SHORT_FILE | F_PID | F_GID | F_COLOR)` to control what's printed

### API on godoc.org

https://godoc.org/github.com/subchen/go-log

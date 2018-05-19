/*
Package log is a simple and configurable Logging in Go, with level, formatters and writers.

It is completely API compatible with the standard library logger.

The simplest way to use log is simply the package-level exported logger:

package main

    import (
    	"os"
    	"github.com/subchen/go-log"
    )

    func main() {
    	log.Print("some message")
    	log.Infof("$HOME = %v", os.Getenv("HOME"))
    	log.Errorln("Got err:", os.ErrPermission)
    }

Output:

    07:34:23.039 INFO some message
    07:34:23.039 INFO $HOME = /home/subchen
    07:34:23.039 ERROR Got err: permission denied


You also can config `log.Default` or new `log.Logger` to customize formatter and writer.

    package main

    import (
    	"os"
    	"github.com/subchen/go-log"
    	"github.com/subchen/go-log/formatters"
    	"github.com/subchen/go-log/writers"
    )

    func main() {
    	logger := &log.Logger{
    		Level:     log.INFO,
    		Formatter: new(formatters.TextFormatter),
    		Out:       &writers.FixedSizeFileWriter{
    			Name:     "/tmp/test.log",
    			MaxSize:  10 * 1024 * 1024, // 10m
    			MaxCount: 10,
    		},
    	}

    	logger.Info("some message")
    }



Output log in `/tmp/test.log`:

    2018-05-19T07:49:05.979+0000 INFO devbox main 9981 example/main.go:17 some message


For a full guide visit https://github.com/subchen/go-log

*/
package log

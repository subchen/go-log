package log

import (
	"io"
)

type CompositeWriters []io.Writer

func (ws *CompositeWriters) Write(p []byte) (n int, err error) {
	for _, w := range *ws {
		n, err = w.Write(p)
	}
	return
}

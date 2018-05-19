package formatters

import (
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// IsTerminal returns whether is a valid tty for io.Writer
func IsTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

package log

import (
	"fmt"
	"os"
	"path"
	"time"
)

// AlwaysNewFileWriter create new log for every process
type AlwaysNewFileWriter struct {
	Name string
	file *os.File
}

func (w *AlwaysNewFileWriter) Write(p []byte) (n int, err error) {
	if w.file == nil {
		w.openFile()
	}

	return w.file.Write(p)
}

func (w *AlwaysNewFileWriter) openFile() (err error) {
	name := fmt.Sprintf("%s.%s", w.Name, time.Now().Format("20060102_150405"))

	// remove symbol link if exist
	os.Remove(w.Name)

	// create symbol
	err = os.Symlink(path.Base(name), w.Name)
	if err != nil {
		return err
	}

	w.file, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}

	return nil
}

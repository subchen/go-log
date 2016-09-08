package log

import (
	"fmt"
	"math"
	"os"
	"path"
)

// FixedSizeFileWriter create new log if log size exceed
type FixedSizeFileWriter struct {
	Name     string
	MaxSize  int64
	MaxCount int

	file        *os.File
	currentSize int64
}

func (w *FixedSizeFileWriter) Write(p []byte) (n int, err error) {
	if w.file == nil {
		w.openCurrentFile()
	} else if w.currentSize > w.MaxSize {
		w.file.Close()
		w.openNextFile()
	}

	w.currentSize += int64(len(p))

	return w.file.Write(p)
}

func (w *FixedSizeFileWriter) openCurrentFile() error {
	name, err := os.Readlink(w.Name)
	if err != nil {
		name = w.getAvailableFileName()

		// create a symlink
		err = os.Symlink(path.Base(name), w.Name)
		if err != nil {
			return err
		}
	} else {
		// convert to abs path
		name = path.Join(path.Dir(w.Name), name)
	}

	w.file, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	stat, err := w.file.Stat()
	if err != nil {
		return err
	}
	w.currentSize = stat.Size()

	return nil
}

func (w *FixedSizeFileWriter) openNextFile() (err error) {
	name := w.getAvailableFileName()

	// remove symbol link
	err = os.Remove(w.Name)
	if err != nil {
		return err
	}

	// create symbol
	err = os.Symlink(path.Base(name), w.Name)
	if err != nil {
		return err
	}

	w.file, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	w.currentSize = 0

	return nil
}

// get available file or oldest file
func (w *FixedSizeFileWriter) getAvailableFileName() string {
	var oldestTime int64 = math.MaxInt64
	var oldestName string

	for i := 0; i < w.MaxCount; i++ {
		name := fmt.Sprintf("%s.%d", w.Name, i)
		stat, err := os.Stat(name)
		if err != nil {
			return name
		}

		if fTime := stat.ModTime().Unix(); fTime < oldestTime {
			oldestTime = fTime
			oldestName = name
		}
	}

	return oldestName
}

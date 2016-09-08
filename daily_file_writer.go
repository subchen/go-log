package log

import (
	"fmt"
	"os"
	"path"
	"time"
)

// DailyFileWriter create new log for every day
type DailyFileWriter struct {
	Name        string
	file        *os.File
	nextDayTime int64
}

func (w *DailyFileWriter) Write(p []byte) (n int, err error) {
	now := time.Now()

	if w.file == nil {
		w.openFile(&now)
	} else if now.Unix() >= w.nextDayTime {
		w.file.Close()
		w.openFile(&now)
	}

	return w.file.Write(p)
}

func (w *DailyFileWriter) openFile(now *time.Time) (err error) {
	name := fmt.Sprintf("%s.%s", w.Name, now.Format("2006-01-02"))

	// remove symbol link if exist
	os.Remove(w.Name)

	// create symbol
	err = os.Symlink(path.Base(name), w.Name)
	if err != nil {
		return err
	}

	w.file, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	year, month, day := now.Date()
	w.nextDayTime = time.Date(year, month, day+1, 0, 0, 0, 0, now.Location()).Unix()

	return nil
}

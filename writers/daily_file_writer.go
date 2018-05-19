package writers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// DailyFileWriter create new log for every day
type DailyFileWriter struct {
	Name     string
	MaxCount int

	file        *os.File
	nextDayTime int64
}

// Write implements io.Writer
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
	name := fmt.Sprintf("%s.%s", w.Name, now.Format("20060102"))

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

	if w.MaxCount > 0 {
		go w.cleanFiles()
	}

	return nil
}

// clean old files
func (w *DailyFileWriter) cleanFiles() {
	dir := path.Dir(w.Name)

	fileList, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	prefix := path.Base(w.Name) + "."

	var matches []string
	for _, f := range fileList {
		if !f.IsDir() && strings.HasPrefix(f.Name(), prefix) {
			matches = append(matches, f.Name())
		}
	}

	if len(matches) > w.MaxCount {
		sort.Sort(sort.Reverse(sort.StringSlice(matches)))

		for _, f := range matches[w.MaxCount:] {
			file := filepath.Join(dir, f)
			os.Remove(file)
		}
	}
}

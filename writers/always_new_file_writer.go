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

// AlwaysNewFileWriter create new log for every process
type AlwaysNewFileWriter struct {
	Name     string
	MaxCount int

	file *os.File
}

// Write implements io.Writer
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

	if w.MaxCount > 0 {
		go w.cleanFiles()
	}

	return nil
}

// clean old files
func (w *AlwaysNewFileWriter) cleanFiles() {
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
		fmt.Println(matches)

		for _, f := range matches[w.MaxCount:] {
			file := filepath.Join(dir, f)
			os.Remove(file)
		}
	}
}

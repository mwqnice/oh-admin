package writer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileWriter struct {
	fileName string
	mutex    sync.RWMutex
	fdm      *FDManager
}

func NewFileWriter(filename string) (*FileWriter, error) {
	fw := &FileWriter{fdm: NewFDManager(), fileName: filename}
	err := fw.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	return fw, nil
}

// OpenFile 设置文件名, 打开并且放入map等待使用
func (fw *FileWriter) OpenFile(filename string) error {
	fw.fileName = filename
	fdValue, _ := fw.fdm.GetFDValue(filename)
	if fdValue != nil {
		return nil
	}

	if err := fw.TryMkdir(); err != nil {
		return err
	}
	err := fw.Open()
	if err != nil {
		return err
	}

	return nil
}

// Open 打开文件
func (fw *FileWriter) Open() error {
	fh, err := os.OpenFile(fw.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	fw.fdm.AddFDMap(fw.fileName, &FDMapValue{fh, time.Now()})

	return nil
}

func (fw *FileWriter) Write(p []byte) (int, error) {
	lens, err := fw.write(p)
	if err != nil {
		return 0, err
	}

	return lens, nil
}

func (fw *FileWriter) TryMkdir() error {
	dirname := filepath.Dir(fw.fileName)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		return errors.New(fmt.Sprintf("failed to create directory %s: %s", dirname, err.Error()))
	}

	return nil
}

func (fw *FileWriter) write(p []byte) (int, error) {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()

	fdValue, err := fw.fdm.GetFDValue(fw.fileName)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("failed to open file %s: %s", fw.fileName, err.Error()))
	}
	lens, err := fdValue.file.Write(p)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("write content fail %s : %s", fw.fileName, err.Error()))
	}

	return lens, nil
}

package writer

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

var fdmStarted = false
var fdm *FDManager

const DefaultFDMaxAge = time.Hour
const DefaultClearInterval = time.Minute * 5

type FDManager struct {
	fileFDMaps sync.Map
	options    *Options
}
type FDMapValue struct {
	file  *os.File
	ctime time.Time
}

type Options struct {
	FDMaxAge      time.Duration
	ClearInterval time.Duration
}
type Option func(options *Options)

func NewFDManager(options ...Option) *FDManager {
	if fdmStarted && fdm != nil {
		return fdm
	}

	fdm = &FDManager{options: NewOptions(options...)}
	fdmStarted = true

	fdm.ClearFD()
	return fdm
}
func NewOptions(options ...Option) *Options {
	opts := &Options{}
	for _, opt := range options {
		opt(opts)
	}

	if opts.FDMaxAge < time.Minute {
		opts.FDMaxAge = DefaultFDMaxAge
	}

	if opts.ClearInterval < time.Minute {
		opts.ClearInterval = DefaultClearInterval
	}

	return opts
}

func WithMaxAge(maxAge time.Duration) Option {
	return func(options *Options) {
		options.FDMaxAge = maxAge
	}
}
func WithClearInterval(clearInterval time.Duration) Option {
	return func(options *Options) {
		options.ClearInterval = clearInterval
	}
}

func (fdm *FDManager) AddFDMap(filename string, value *FDMapValue) {
	fdm.fileFDMaps.Store(filename, value)
}
func (fdm *FDManager) RemoveFDMap(filename string) {
	fdm.fileFDMaps.Delete(filename)
}
func (fdm *FDManager) GetFDValue(filename string) (*FDMapValue, error) {
	fdValue, ok := fdm.fileFDMaps.Load(filename)
	if !ok {
		return nil, errors.New("fdValue empty")
	}

	return fdValue.(*FDMapValue), nil
}

// ClearFD 关闭一段时间未写入的文件
func (fdm *FDManager) ClearFD() {
	go func() {
		for {
			fdm.fileFDMaps.Range(func(key, value interface{}) bool {
				info := value.(*FDMapValue)
				fileName := key.(string)
				fi, _ := info.file.Stat()
				// 超过x个小时未修改 则关闭文件
				if fi.ModTime().Before(time.Now().Add(-fdm.options.FDMaxAge)) {
					fdm.fileFDMaps.Delete(fileName)
					err := info.file.Close()
					if err != nil {
						println(fmt.Sprintf("close file fail : %s : %s", fileName, err.Error()))
					}
				}
				return true
			})

			time.Sleep(fdm.options.ClearInterval)
		}
	}()

}

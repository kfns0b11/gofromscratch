package cuslog

import (
	"io"
	"sync"
	"unsafe"
)

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

var stdlogger = New()

func New(opts ...Option) *logger {
	logger := &logger{opt: initOptions(opts...)}
	logger.entryPool = &sync.Pool{New: func() interface{} { return entry(logger) }}
	return logger
}

func StdLogger() *logger {
	return stdlogger
}

func SetOptions(opts ...Option) {
	stdlogger.SetOptions(opts...)
}

func (l *logger) SetOptions(opts ...Option) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}

func (l *logger) Write(data []byte) (int, error) {
	l.entry().write(l.opt.stdlevel, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func Writer() io.Writer {
	return stdlogger
}

func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

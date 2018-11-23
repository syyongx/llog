package handler

import (
	"bufio"
	"github.com/kataras/iris/core/errors"
	"github.com/syyongx/llog/types"
	"os"
	"sync"
	"time"
)

type FlushMode uint8

const (
	FlushModeTicker = iota + 1
	FlushModeLimit
)

// File handler.
type File struct {
	Processing

	Path     string
	FilePerm os.FileMode
	Fd       *os.File
	Bufio
}

type Bufio struct {
	sync.Mutex

	useBufio      bool
	bufioSize     int
	flushMode     FlushMode
	flushInterval time.Duration
	ioWriter      *bufio.Writer
}

// New file handler
func NewFile(path string, filePerm os.FileMode, level int, bubble bool) *File {
	file := &File{
		Path:     path,
		FilePerm: filePerm,
	}
	file.SetLevel(level)
	file.SetBubble(bubble)
	file.Writer = file.Write

	return file
}

// Set flush config.
func (f *File) SetBufio(size int, mode FlushMode, interval time.Duration) error {
	if size < 0 {
		return errors.New("size invalid")
	}
	f.useBufio = true
	f.bufioSize = size
	f.flushMode = mode
	if mode == FlushModeTicker {
		if interval < time.Second {
			interval = time.Second
		}
		f.flushInterval = interval
		// async auto flush bufio
		go f.tickerFlush()
	}

	return nil
}

// Write to file.
func (f *File) Write(record *types.Record) {
	f.Lock()
	defer f.Unlock()

	if f.Fd == nil {
		fd, err := os.OpenFile(f.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, f.FilePerm)
		if err != nil {
			// ...
		}
		f.Fd = fd
		// use bufio
		if f.useBufio {
			if f.ioWriter == nil {
				f.ioWriter = bufio.NewWriterSize(f.Fd, f.bufioSize)
			} else {
				f.ioWriter.Reset(f.Fd)
			}
		}
	}
	// write
	var err error
	b := record.Formatted.Bytes()
	if f.useBufio {
		_, err = f.ioWriter.Write(b)
	} else {
		_, err = f.Fd.Write(b)
	}
	if err != nil {
		//...
	}
}

// flush
func (f *File) Flush() (err error) {
	if !f.useBufio {
		return
	}

	f.Lock()
	err = f.ioWriter.Flush()
	f.Unlock()
	return
}

// Close writer
func (f *File) Close() {
	if f.useBufio && f.ioWriter != nil {
		f.ioWriter.Flush()
	}
	f.Fd.Close()
	f.Fd = nil
}

// Auto flush
func (f *File) tickerFlush() {
	ticker := time.NewTicker(f.flushInterval)
	for range ticker.C {
		f.Flush()
	}
}

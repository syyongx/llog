package handler

import (
	"bufio"
	"github.com/syyongx/llog/types"
	"os"
)

// File handler.
type File struct {
	Processing

	Path     string
	FilePerm os.FileMode
	Fd       *os.File
	ioWriter *bufio.Writer
}

// New file handler
func NewFile(path string, level int, bubble bool, filePerm os.FileMode) *File {
	file := &File{
		Path:     path,
		FilePerm: filePerm,
	}
	file.SetLevel(level)
	file.SetBubble(bubble)
	file.Writer = file.Write
	return file
}

// Write to file.
func (f *File) Write(record *types.Record) {
	if f.Fd == nil {
		fd, err := os.OpenFile(f.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, f.FilePerm)
		if err != nil {
			// ...
		}
		f.Fd = fd
		f.ioWriter = bufio.NewWriterSize(f.Fd, 1)
	}
	_, err := f.Fd.Write(record.Formatted.Bytes())
	if err != nil {
		//...
	}
}

// Close writer
func (f *File) Close() {
	f.Fd.Close()
	f.Fd = nil
}

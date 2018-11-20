package handler

import (
	"github.com/syyongx/llog/types"
	"os"
)

// File handler.
type File struct {
	Processing

	Path     string
	FilePerm os.FileMode
	Fd       *os.File
}

// New file handler
func NewFile(path string, level int, bubble bool, filePerm os.FileMode) *File {
	file := &File{
		Path:     path,
		FilePerm: filePerm,
	}
	file.SetLevel(level)
	file.SetBubble(bubble)
	return file
}

// Handles a record.
func (f *File) Handle(record *types.Record) bool {
	if !f.IsHandling(record) {
		return false
	}
	if f.processors != nil {
		f.ProcessRecord(record)
	}
	err := f.GetFormatter().Format(record)
	if err != nil {
		return false
	}
	f.Write(record)

	return false == f.GetBubble()
}

// Handles a set of records.
func (f *File) HandleBatch(records []*types.Record) {
	for _, record := range records {
		f.Handle(record)
	}
}

// Write to file.
func (f *File) Write(record *types.Record) {
	if f.Fd == nil {
		fd, err := os.OpenFile(f.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, f.FilePerm)
		if err != nil {
			// ...
		}
		f.Fd = fd
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

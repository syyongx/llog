package handler

import (
	"os"
	"strings"
	"regexp"
	"github.com/syyongx/llog/types"
)

// Stores logs to files that are rotated every day and a limited number of files are kept.
//
// This rotation is only intended to be used as a workaround. Using logrotate to
// handle the rotation is strongly encouraged when you can use it.
type RotatingFile struct {
	f File

	filename       string
	maxFiles       int
	mustRotate     bool
	nextRotation   string
	filenameFormat string
	dateFormat     string
}

// level: The minimum logging level at which this handler will be triggered
// bubble: Whether the messages that are handled can bubble up the stack or not
// filePerm: Optional file permissions (default (0644) are only for owner read/write)
func NewRotatingFile(filename string, maxFiles, level int, bubble bool, filePerm os.FileMode) *RotatingFile {
	rf := &RotatingFile{
		filename:       filename,
		maxFiles:       maxFiles,
		filenameFormat: "{filename}-{date}",
		dateFormat:     "2016-01-02",
	}
	rf.f.SetLevel(level)
	rf.f.SetBubble(bubble)
	return rf
}

func (rf *RotatingFile) SetFilenameFormat(filenameFormat, dateFormat string) {
	match, _ := regexp.MatchString("^2006(([/_.-]?01)([/_.-]?02)?)?$", dateFormat)
	if !match {
		// error
		return
	}
	if n := strings.Index(filenameFormat, "{date}"); n < 0 {
		// error
		return
	}

	rf.filenameFormat = filenameFormat
	rf.dateFormat = dateFormat
}

func (rf *RotatingFile) Write(record *types.Record) {

}

func (rf *RotatingFile) Close() {

}

// Roate
func (rf *RotatingFile) rotate() {

}

// Get timed filename
func (rf *RotatingFile) getTimedFilename() {

}

// Get blob pattern
func (rf *RotatingFile) getGlobPattern() {

}

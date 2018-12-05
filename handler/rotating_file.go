package handler

import (
	"errors"
	"github.com/syyongx/llog/types"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// some rotating date formats
const (
	FilePerDay   string = "2006-01-02"
	FilePerMonth string = "2006-01"
	FilePerYear  string = "2006"
)

// RotatingFile Stores logs to files that are rotated every day and a limited number of files are kept.
//
// This rotation is only intended to be used as a workaround. Using logrotate to
// handle the rotation is strongly encouraged when you can use it.
type RotatingFile struct {
	*File

	filename       string
	maxFiles       int
	mustRotate     bool
	nextRotation   int
	filenameFormat string
	dateFormat     string
	sync.Mutex
}

// NewRotatingFile New rotatingFile handler
// level: The minimum logging level at which this handler will be triggered
// bubble: Whether the messages that are handled can bubble up the stack or not
// filePerm: Optional file permissions (default (0644) are only for owner read/write)
func NewRotatingFile(filename string, filePerm os.FileMode, maxFiles, level int, bubble bool) *RotatingFile {
	rf := &RotatingFile{
		filename:       filename,
		maxFiles:       maxFiles,
		filenameFormat: "{filename}-{date}",
		dateFormat:     FilePerDay,
	}
	rf.nextRotation = rf.day(time.Now().AddDate(0, 0, 1))
	path := rf.timedFilename()
	rf.File = NewFile(path, filePerm, level, bubble)
	rf.File.Writer = rf.Write
	return rf
}

// SetFilenameFormat Set filename format.
func (rf *RotatingFile) SetFilenameFormat(filenameFormat, dateFormat string) error {
	// validate data format
	match, _ := regexp.MatchString("^2006(([/_.-]?01)([/_.-]?02)?)?$", dateFormat)
	if !match {
		return errors.New("invalid date format")
	}
	if n := strings.Index(filenameFormat, "{date}"); n < 0 {
		return errors.New("invalid filename format, format should contain at least {date}")
	}

	rf.filenameFormat = filenameFormat
	rf.dateFormat = dateFormat
	rf.Path = rf.timedFilename()
	rf.Close()

	return nil
}

// Write to file.
func (rf *RotatingFile) Write(record *types.Record) {
	// need rotate
	rf.Lock()
	if rf.nextRotation < rf.day(record.Datetime) {
		rf.mustRotate = true
		rf.Close()
	}
	rf.Unlock()

	rf.File.Write(record)
}

// Close the handler.
func (rf *RotatingFile) Close() {
	rf.File.Close()

	if rf.mustRotate {
		// do ratate
		rf.rotate()
	}
}

// Rotates the files.
func (rf *RotatingFile) rotate() error {
	// update path
	rf.Path = rf.timedFilename()
	rf.Fd = nil
	// tomorrow
	rf.nextRotation = rf.day(time.Now().AddDate(0, 0, 1))
	// skip remove old files if files are unlimited
	if rf.maxFiles == 0 {
		return nil
	}
	// async remove old files.
	go rf.removeOldLogs()

	rf.mustRotate = false
	return nil
}

// Get timed filename
func (rf *RotatingFile) timedFilename() string {
	dir := filepath.Dir(rf.filename)
	basename := filepath.Base(rf.filename)
	ext := filepath.Ext(rf.filename)
	if ext != "" {
		basename = basename[:strings.Index(basename, ext)]
	}

	date := time.Unix(time.Now().Unix(), 0).Format(rf.dateFormat)
	timedFilename := strings.NewReplacer("{filename}", basename, "{date}", date).Replace(dir + "/" + rf.filenameFormat)
	timedFilename += ext

	return timedFilename
}

// Get blob pattern
func (rf *RotatingFile) globPattern() string {
	dir := filepath.Dir(rf.filename)
	basename := filepath.Base(rf.filename)
	ext := filepath.Ext(rf.filename)
	if ext != "" {
		basename = basename[:strings.Index(basename, ext)]
	}

	glob := strings.NewReplacer("{filename}", basename, "{date}", "*").Replace(dir + "/" + rf.filenameFormat)
	glob += ext

	return glob
}

// get tomorrow day
func (rf *RotatingFile) day(t time.Time) int {
	day, _ := strconv.Atoi(t.Format("20060102"))
	return day
}

// Remove old logs.
func (rf *RotatingFile) removeOldLogs() {
	files, err := filepath.Glob(rf.globPattern())
	if err != nil {
		return
	}
	if len(files) <= rf.maxFiles {
		return
	}
	// Sorting the files by name to remove the older ones
	sort.Strings(files)
	for _, file := range files[:len(files)-rf.maxFiles] {
		os.Remove(file)
	}
}

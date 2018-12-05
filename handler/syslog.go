package handler

import (
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/types"
	"log/syslog"
)

// Syslog struct definition
type Syslog struct {
	Processing

	SysWriter *syslog.Writer
}

// NewSyslog New establishes a new connection to the system log daemon. Each
// write to the returned writer sends a log message with the given
// priority (a combination of the syslog facility and severity) and
// prefix tag. If tag is empty, the os.Args[0] is used.
func NewSyslog(priority syslog.Priority, tag string, level int, bubble bool) (*Syslog, error) {
	sys := &Syslog{}
	w, err := syslog.New(priority, tag)
	if err != nil {
		return nil, err
	}
	sys.SysWriter = w
	sys.SetLevel(level)
	sys.SetBubble(bubble)
	sys.SetFormatter(sys.GetDefaultFormatter())
	sys.Writer = sys.Write
	return sys, nil
}

// Write to console.
func (s *Syslog) Write(record *types.Record) {
	if s.SysWriter == nil {
		return
	}
	var fn func(m string) error
	switch record.Level {
	case types.DEBUG:
		fn = s.SysWriter.Debug
	case types.INFO:
		fn = s.SysWriter.Info
	case types.NOTICE:
		fn = s.SysWriter.Notice
	case types.WARNING:
		fn = s.SysWriter.Warning
	case types.ERROR:
		fn = s.SysWriter.Err
	case types.CRITICAL:
		fn = s.SysWriter.Crit
	case types.ALERT:
		fn = s.SysWriter.Alert
	case types.EMERGENCY:
		fn = s.SysWriter.Emerg
	}
	fn(record.Formatted.String())
}

// GetDefaultFormatter Gets the default syslog formatter.
func (s *Syslog) GetDefaultFormatter() types.Formatter {
	return formatter.NewLine("%Channel%.%LevelName%: %Message% %Context% %Extra%", "")
}

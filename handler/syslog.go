package handler

import (
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/types"
	"log/syslog"
)

type Syslog struct {
	Processing

	Writer *syslog.Writer
}

// New establishes a new connection to the system log daemon. Each
// write to the returned writer sends a log message with the given
// priority (a combination of the syslog facility and severity) and
// prefix tag. If tag is empty, the os.Args[0] is used.
func NewSyslog(priority syslog.Priority, tag string, level int, bubble bool) *Syslog {
	sys := &Syslog{}
	w, err := syslog.New(priority, tag)
	if err != nil {
		// ...
	}
	sys.Writer = w
	sys.SetLevel(level)
	sys.SetBubble(bubble)
	sys.SetFormatter(sys.GetDefaultFormatter())
	return sys
}

// Write to console.
func (s *Syslog) Write(record *types.Record) {
	if s.Writer == nil {
		return
	}
	var fn func(m string) error
	switch record.Level {
	case types.DEBUG:
		fn = s.Writer.Debug
	case types.INFO:
		fn = s.Writer.Info
	case types.NOTICE:
		fn = s.Writer.Notice
	case types.WARNING:
		fn = s.Writer.Warning
	case types.ERROR:
		fn = s.Writer.Err
	case types.CRITICAL:
		fn = s.Writer.Crit
	case types.ALERT:
		fn = s.Writer.Alert
	case types.EMERGENCY:
		fn = s.Writer.Emerg
	}
	fn(record.Formatted.String())
}

// Gets the default syslog formatter.
func (s *Syslog) GetDefaultFormatter() types.Formatter {
	return formatter.NewLine("%Channel%.%LevelName%: %Message% %Context% %Extra%", "")
}

// Package llog is a log tool library. php Monolog implementation in Go.
// Source code and other details for the project are available at GitHub:
//
//	https://github.com/syyongx/llog
//
package llog

import (
	"errors"
	"fmt"
	"github.com/syyongx/llog/types"
	"strings"
	"time"
)

// Logger logger struct
type Logger struct {
	name       string
	levels     map[int]string
	handlers   []types.IHandler
	processors []types.Processor
	timezone   string
}

var levels = map[int]string{
	types.DEBUG:     "debug",
	types.INFO:      "info",
	types.NOTICE:    "notice",
	types.WARNING:   "warning",
	types.ERROR:     "error",
	types.CRITICAL:  "critical",
	types.ALERT:     "alert",
	types.EMERGENCY: "emergency",
}

// NewLogger new logger
func NewLogger(name string) *Logger {
	return &Logger{
		name:   name,
		levels: levels,
	}
}

// GetName get name
func (l *Logger) GetName() string {
	return l.name
}

// PushHandler pushes a handler on to the stack.
func (l *Logger) PushHandler(h types.IHandler) {
	l.handlers = append([]types.IHandler{h}, l.handlers...)
}

// PopHandler Pops a handler from the stack.
func (l *Logger) PopHandler() (types.IHandler, error) {
	if len(l.handlers) < 1 {
		return nil, errors.New("you tried to pop from an empty handler slice")
	}
	first := l.handlers[0]
	l.handlers = l.handlers[1:]
	return first, nil
}

// SetHandlers Set handlers, replacing all existing ones.
func (l *Logger) SetHandlers(handlers []types.IHandler) {
	l.handlers = l.handlers[:0]
	for i := len(handlers); i > 0; i-- {
		l.PushHandler(handlers[i-1])
	}
}

// GetHandlers Get handlers
func (l *Logger) GetHandlers() []types.IHandler {
	return l.handlers
}

// PushProcessor Pushes a processor on to the stack.
func (l *Logger) PushProcessor(p types.Processor) {
	l.processors = append([]types.Processor{p}, l.processors...)
}

// PopProcessor Pops a processor from the stack.
func (l *Logger) PopProcessor() (types.Processor, error) {
	if len(l.processors) < 1 {
		return nil, errors.New("you tried to pop from an empty processor stack")
	}
	first := l.processors[0]
	l.processors = l.processors[1:]
	return first, nil
}

// GetProcessor Get processors
func (l *Logger) GetProcessor() []types.Processor {
	return l.processors
}

// AddRecord Adds a log record.
func (l *Logger) AddRecord(level int, message string) bool {
	hKey := -1
	record := types.GetRecord()
	defer types.ReleaseRecord(record)
	record.Level = level
	for i, v := range l.handlers {
		if v.IsHandling(record) {
			hKey = i
		}
	}
	if hKey == -1 {
		return false
	}
	levelName, err := l.GetLevelName(level)
	if err != nil {
		// ignore errors
		return false
	}
	record.Message = message
	record.LevelName = levelName
	record.Channel = l.name
	record.Datetime = time.Now()

	for _, p := range l.processors {
		p(record)
	}
	for j, h := range l.handlers {
		if hKey < j {
			continue
		}
		if h.Handle(record) { // Will not bubble
			break
		}
	}

	return true
}

// GetLevels Gets all supported logging levels.
func (l *Logger) GetLevels() map[int]string {
	return l.levels
}

// GetLevelName Gets the name of the logging level.
func (l *Logger) GetLevelName(level int) (string, error) {
	if v, ok := l.levels[level]; ok {
		return v, nil
	}
	return "", errors.New("level is not defined")
}

// GetLevelByName Gets the value by the logging level name.
func (l *Logger) GetLevelByName(levelName string) (int, error) {
	levelName = strings.ToLower(levelName)

	for val, name := range l.levels {
		if name == levelName {
			return val, nil
		}
	}

	return 0, errors.New("level is not defined")
}

// IsHandling Checks whether the Logger has a handler that listens on the given level.
func (l *Logger) IsHandling(level int) bool {
	return true
}

// Log Logs with an arbitrary level.
func (l *Logger) Log(level int, message string) {
	if _, ok := l.levels[level]; !ok {
		return
	}

	l.AddRecord(level, message)
}

// Debug Detailed debug information.
func (l *Logger) Debug(message interface{}) {
	l.AddRecord(types.DEBUG, l.String(message))
}

// Info Interesting events.
// Example: User logs in, SQL logs.
func (l *Logger) Info(message interface{}) {
	l.AddRecord(types.INFO, l.String(message))
}

// Notice Normal but significant events.
func (l *Logger) Notice(message interface{}) {
	l.AddRecord(types.NOTICE, l.String(message))
}

// Warning Exceptional occurrences that are not errors.
// Example: Use of deprecated APIs, poor use of an API, undesirable things
// that are not necessarily wrong.
func (l *Logger) Warning(message interface{}) {
	l.AddRecord(types.WARNING, l.String(message))
}

// Error Runtime errors that do not require immediate action but should typically
// be logged and monitored.
func (l *Logger) Error(message interface{}) {
	l.AddRecord(types.ERROR, l.String(message))
}

// Alert Adds a log record at the ALERT level.
func (l *Logger) Alert(message interface{}) {
	l.AddRecord(types.ALERT, l.String(message))
}

// Emergency System is unusable.
func (l *Logger) Emergency(message interface{}) {
	l.AddRecord(types.EMERGENCY, l.String(message))
}

// SetTimezone Set the timezone to be used for the timestamp of log records.
func (l *Logger) SetTimezone(tz string) {
	l.timezone = tz
}

// GetTimezone Get timezone
func (l *Logger) GetTimezone(tz string) string {
	return l.timezone
}

// String Stringify
func (l *Logger) String(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case []byte:
		return string(data.([]byte))
	}
	return fmt.Sprint(data)
}

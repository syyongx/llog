package llog

import (
	"errors"
	"fmt"
	"github.com/syyongx/llog/handler"
	"github.com/syyongx/llog/types"
	"github.com/syyongx/llog/processor"
	"time"
)

// logger struct
type Logger struct {
	name       string
	levels     map[int]string
	handlers   []handler.IHandler
	processors []processor.Processor
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

// new logger
func NewLogger(name string) *Logger {
	return &Logger{
		name:   name,
		levels: levels,
	}
}

func (l *Logger) GetName() string {
	return l.name
}

// Pushes a handler on to the stack.
func (l *Logger) PushHandler(h handler.IHandler) {
	l.handlers = append([]handler.IHandler{h}, l.handlers...)
}

// Pops a handler from the stack.
func (l *Logger) PopHandler() (handler.IHandler, error) {
	if len(l.handlers) < 1 {
		return nil, errors.New("You tried to pop from an empty handler slice.")
	}
	first := l.handlers[0]
	l.handlers = l.handlers[1:]
	return first, nil
}

// Set handlers, replacing all existing ones.
func (l *Logger) SetHandlers(handlers []handler.IHandler) {
	l.handlers = l.handlers[:0]
	for i := len(handlers); i > 0; i-- {
		l.PushHandler(handlers[i-1])
	}
}

// Get handlers
func (l *Logger) GetHandlers() []handler.IHandler {
	return l.handlers
}

// Pushes a processor on to the stack.
func (l *Logger) PushProcessor(p processor.Processor) {
	l.processors = append([]processor.Processor{p}, l.processors...)
}

// Pops a processor from the stack.
func (l *Logger) PopProcessor() (processor.Processor, error) {
	if len(l.processors) < 1 {
		return nil, errors.New("You tried to pop from an empty processor slice.")
	}
	first := l.processors[0]
	l.processors = l.processors[1:]
	return first, nil
}

// Get processors
func (l *Logger) GetProcessor() []processor.Processor {
	return l.processors
}

// Adds a log record.
func (l *Logger) AddRecord(level int, message string) (bool, error) {
	hKey := -1
	record := &types.Record{Level: level}
	for i, v := range l.handlers {
		if v.IsHandling(record) {
			hKey = i
		}
	}
	if hKey == -1 {
		return false, nil
	}
	levelName, err := l.GetLevelName(level)
	if err != nil {
		return false, err
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
		if h.Handle(record) {
			break
		}
	}

	return true, nil
}

// Gets all supported logging levels.
func (l *Logger) GetLevels() map[int]string {
	return l.levels
}

// Gets the name of the logging level.
func (l *Logger) GetLevelName(level int) (string, error) {
	if v, ok := l.levels[level]; ok {
		return v, nil
	}
	return "", errors.New("Level is not defined")
}

// Checks whether the Logger has a handler that listens on the given level.
func (l *Logger) IsHandling(level int) bool {
	return true
}

// Adds a log record at an arbitrary level.
func (l *Logger) Log(level int, message string) {
	if _, ok := l.levels[level]; !ok {
		//
	}
	l.AddRecord(level, message)
}

// Adds a log record at the DEBUG level.
func (l *Logger) Debug(message interface{}) {
	l.AddRecord(types.DEBUG, l.String(message))
}

// Adds a log record at the INFO level.
func (l *Logger) Info(message interface{}) {
	l.AddRecord(types.INFO, l.String(message))
}

// Adds a log record at the NOTICE level.
func (l *Logger) Notice(message interface{}) {
	l.AddRecord(types.NOTICE, l.String(message))
}

// Adds a log record at the WARNING level.
func (l *Logger) Warning(message interface{}) {
	l.AddRecord(types.WARNING, l.String(message))
}

// Adds a log record at the ERROR level.
func (l *Logger) Error(message interface{}) {
	l.AddRecord(types.ERROR, l.String(message))
}

// Adds a log record at the ALERT level.
func (l *Logger) Alert(message interface{}) {
	l.AddRecord(types.ALERT, l.String(message))
}

// Adds a log record at the EMERGENCY level.
func (l *Logger) Emergency(message interface{}) {
	l.AddRecord(types.EMERGENCY, l.String(message))
}

// Set the timezone to be used for the timestamp of log records.
func (l *Logger) SetTimezone(tz string) {
	l.timezone = tz
}

// Get timezone
func (l *Logger) GetTimezone(tz string) string {
	return l.timezone
}

// Stringfy
func (l *Logger) String(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case []byte:
		return string(data.([]byte))
	}
	return fmt.Sprint(data)
}

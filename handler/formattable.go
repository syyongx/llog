package handler

import (
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/types"
	"time"
)

// Formattable struct definition
type Formattable struct {
	formatter types.Formatter
}

// SetFormatter Set formatter
func (f *Formattable) SetFormatter(formatter types.Formatter) {
	f.formatter = formatter
}

// GetFormatter Get formatter
func (f *Formattable) GetFormatter() types.Formatter {
	return f.formatter
}

// GetDefaultFormatter Gets the default formatter.
func (f *Formattable) GetDefaultFormatter() types.Formatter {
	return formatter.NewLine(formatter.DefaultFormat, time.RFC3339)
}

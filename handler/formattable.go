package handler

import (
	"github.com/syyongx/llog/formatter"
	"time"
	"github.com/syyongx/llog/types"
)

type Formattable struct {
	formatter types.IFormatter
}

// Set formatter
func (f *Formattable) SetFormatter(formatter types.IFormatter) {
	f.formatter = formatter
}

// Get formatter
func (f *Formattable) GetFormatter() types.IFormatter {
	return f.formatter
}

//Gets the default formatter.
func (f *Formattable) GetDefaultFormatter() types.IFormatter {
	return formatter.NewLine(formatter.DefaultFormat, time.RFC3339)
}

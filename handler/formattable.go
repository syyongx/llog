package handler

import (
	"github.com/syyongx/llog/formatter"
	"time"
	"github.com/syyongx/llog/types"
)

type Formattable struct {
	formatter types.Formatter
}

// Set formatter
func (f *Formattable) SetFormatter(formatter types.Formatter) {
	f.formatter = formatter
}

// Get formatter
func (f *Formattable) GetFormatter() types.Formatter {
	return f.formatter
}

//Gets the default formatter.
func (f *Formattable) GetDefaultFormatter() types.Formatter {
	return formatter.NewLine(formatter.DefaultFormat, time.RFC3339)
}

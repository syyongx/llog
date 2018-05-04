package handler

import "github.com/syyongx/llog/formatter"

type Formattable struct {
	formatter formatter.IFormatter
}

// Set formatter
func (f *Formattable) SetFormatter(formatter formatter.IFormatter) {
	f.formatter = formatter
}

// Get formatter
func (f *Formattable) GetFormatter() formatter.IFormatter {
	return f.formatter
}

// Gets the default formatter.
//func (f *Formattable) GetDefaultFormatter() formatter.IFormatter {
//	//return formatter.NewLine()
//}

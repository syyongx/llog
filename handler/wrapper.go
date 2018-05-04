package handler

import (
	"github.com/syyongx/llog/types"
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/processor"
)

//
type Wrapper struct {
	handler IHandler
}

// new handler wrapper
func NewWrapper(handler IHandler) *Wrapper {
	return &Wrapper{
		handler: handler,
	}
}

//
func (w *Wrapper) IsHandling(record types.Record) bool {
	return true
}

func (w *Wrapper) Handle(record types.Record) bool {
	return w.handler.Handle(record)
}

func (w *Wrapper) HandleBatch(records []types.Record) {
}

func (w *Wrapper) Close(record types.Record) {

}

func (w *Wrapper) PushProcessor(processor processor.Processor) *IHandler {
	return nil
}

//func (w *Wrapper) PopProcessor() processor.Processor {
//
//}

func (w *Wrapper) SetFormatter(formatter *formatter.IFormatter) {

}

func (w *Wrapper) GetFormatter() *formatter.IFormatter {
	return nil
}

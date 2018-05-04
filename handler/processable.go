package handler

import (
	"github.com/syyongx/llog/processor"
	"github.com/syyongx/llog/types"
)

type Processable struct {
	processors []processor.Processor
}

// push processor
func (p *Processable) PushProcessor(processor.Processor) {

}

// pop processor
func (p *Processable) PopProcessor() processor.Processor {
	return nil
}

// Processes a record.
func (p *Processable) ProcessRecord(record *types.Record) *types.Record {
	for _, v := range p.processors {
		record = v(record)
	}

	return record
}

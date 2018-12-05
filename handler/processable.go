package handler

import (
	"github.com/syyongx/llog/types"
)

// Processable struct definition
type Processable struct {
	processors []types.Processor
}

// PushProcessor push processor
func (p *Processable) PushProcessor(types.Processor) {

}

// PopProcessor pop processor
func (p *Processable) PopProcessor() types.Processor {
	return nil
}

// ProcessRecord Processes a record.
func (p *Processable) ProcessRecord(record *types.Record) {
	for _, v := range p.processors {
		v(record)
	}
}

package handler

import (
	"github.com/syyongx/llog/types"
)

type Processable struct {
	processors []types.Processor
}

// push processor
func (p *Processable) PushProcessor(types.Processor) {

}

// pop processor
func (p *Processable) PopProcessor() types.Processor {
	return nil
}

// Processes a record.
func (p *Processable) ProcessRecord(record *types.Record){
	for _, v := range p.processors {
		v(record)
	}
}

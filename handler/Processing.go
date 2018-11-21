package handler

import "github.com/syyongx/llog/types"

type Processing struct {
	Handler
	Processable
	Formattable

	Writer func(*types.Record)
}

// Handles a record.
func (p *Processing) Handle(record *types.Record) bool {
	if !p.IsHandling(record) {
		return false
	}
	if p.processors != nil {
		p.ProcessRecord(record)
	}
	err := p.GetFormatter().Format(record)
	if err != nil {
		return false
	}
	p.Writer(record)

	return false == p.GetBubble()
}

// Handles a set of records.
func (p *Processing) HandleBatch(records []*types.Record) {
	for _, record := range records {
		p.Handle(record)
	}
}

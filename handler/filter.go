package handler

import "github.com/syyongx/llog/types"

type Filter struct {
	Handler
	Processable

	h              types.IHandler
	acceptedLevels map[int]int
}

func NewFilter(handler types.IHandler, minLevels, maxLevels []int, bubble bool) *Filter {
	f := &Filter{
		h: handler,
	}
	f.acceptedLevels = f.SetAcceptedLevels(minLevels, maxLevels)
	f.SetBubble(bubble)

	return f
}

func (f *Filter) IsHandling(record *types.Record) bool {
	_, ok := f.acceptedLevels[record.Level]
	return ok
}

// Handle
func (f *Filter) Handle(record *types.Record) bool {
	if !f.IsHandling(record) {
		return false
	}
	if f.processors != nil {
		f.ProcessRecord(record)
	}
	f.h.Handle(record)

	return false == f.GetBubble()
}

// HandleBatch
func (f *Filter) HandleBatch(records []*types.Record) {
	filtered := make([]*types.Record, 0, len(records))
	for _, record := range records {
		if f.IsHandling(record) {
			filtered = append(filtered, record)
		}
	}
	f.h.HandleBatch(filtered)
}

// Set acceptedLevels
func (f *Filter) SetAcceptedLevels(minLevels, maxLevels []int) map[int]int {
	return nil
}

// Get acceptedLevels
func (f *Filter) GetAcceptedLevels() map[int]int {
	return f.acceptedLevels
}

package handler

import "github.com/syyongx/llog/types"

// Filter struct definition
type Filter struct {
	Handler
	Processable

	h              types.IHandler
	acceptedLevels map[int]int
}

// NewFilter New Filter
func NewFilter(handler types.IHandler, minLevels, maxLevels []int, bubble bool) *Filter {
	filter := &Filter{
		h: handler,
	}
	filter.acceptedLevels = filter.SetAcceptedLevels(minLevels, maxLevels)
	filter.SetBubble(bubble)

	return filter
}

// IsHandling Is Handling
func (f *Filter) IsHandling(record *types.Record) bool {
	_, ok := f.acceptedLevels[record.Level]
	return ok
}

// Handle log record
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

// HandleBatch log records
func (f *Filter) HandleBatch(records []*types.Record) {
	filtered := make([]*types.Record, 0, len(records))
	for _, record := range records {
		if f.IsHandling(record) {
			filtered = append(filtered, record)
		}
	}
	f.h.HandleBatch(filtered)
}

// SetAcceptedLevels Set acceptedLevels
func (f *Filter) SetAcceptedLevels(minLevels, maxLevels []int) map[int]int {
	// TODO: ...
	return nil
}

// GetAcceptedLevels Get acceptedLevels
func (f *Filter) GetAcceptedLevels() map[int]int {
	return f.acceptedLevels
}

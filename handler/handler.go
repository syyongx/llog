package handler

import (
	"github.com/syyongx/llog/types"
)

type Handler struct {
	level  int
	bubble bool
}

// Checks whether the given record will be handled by this handler.
func (h *Handler) IsHandling(record types.Record) bool {
	return record["level"].(int) >= h.level
}

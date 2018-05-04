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

// Set level
func (h *Handler) SetLevel(level int) {
	h.level = level
}

// Get level
func (h *Handler) GetLevel() int {
	return h.level
}

// Set bubble
func (h *Handler) SetBubble(bubble bool) {
	h.bubble = bubble
}

// Get bubble
func (h *Handler) GetBubble() bool {
	return h.bubble
}

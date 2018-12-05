// Package handler provides some commonly used handlers。 such as: file, mail, syslog and more
package handler

import (
	"github.com/syyongx/llog/types"
)

// Handler struct definition
type Handler struct {
	level  int
	bubble bool
}

// IsHandling Checks whether the given record will be handled by this handler.
func (h *Handler) IsHandling(record *types.Record) bool {
	return record.Level >= h.level
}

// SetLevel Set level
func (h *Handler) SetLevel(level int) {
	h.level = level
}

// GetLevel Get level
func (h *Handler) GetLevel() int {
	return h.level
}

// SetBubble Set bubble
func (h *Handler) SetBubble(bubble bool) {
	h.bubble = bubble
}

// GetBubble Get bubble
func (h *Handler) GetBubble() bool {
	return h.bubble
}

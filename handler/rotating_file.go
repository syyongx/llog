package handler

// Stores logs to files that are rotated every day and a limited number of files are kept.
//
// This rotation is only intended to be used as a workaround. Using logrotate to
// handle the rotation is strongly encouraged when you can use it.
type RotatingFile struct {
	File
}

func NewRotatingFile() *RotatingFile {
	return &RotatingFile{}
}

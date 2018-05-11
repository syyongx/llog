package processor

import (
	"time"
	"fmt"
	"github.com/syyongx/llog/types"
)

var Uid = func(record *types.Record) {
	record.Extra["uid"] = generateUid()
}

// Generate uniqid
func generateUid() string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	return fmt.Sprintf("%08x%05x", sec, usec)
}

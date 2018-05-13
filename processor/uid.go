package processor

import (
	"time"
	"fmt"
	"github.com/syyongx/llog/types"
)

// Uniqid processor.
var Uid = func(record *types.Record) {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	record.Extra["Uid"] = fmt.Sprintf("%08x%05x", sec, usec)
}

package processor

import (
	"os"
	"github.com/syyongx/llog/types"
)

// Get process id.
var ProcessId = func(record *types.Record) {
	record.Extra["Pid"] = os.Getpid()
}

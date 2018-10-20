package processor

import (
	"github.com/syyongx/llog/types"
	"os"
)

// Get process id.
var ProcessId = func(record *types.Record) {
	record.Extra["Pid"] = os.Getpid()
}

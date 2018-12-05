package processor

import (
	"github.com/syyongx/llog/types"
	"os"
)

// ProcessID Get process ID.
var ProcessID = func(record *types.Record) {
	record.Extra["Pid"] = os.Getpid()
}

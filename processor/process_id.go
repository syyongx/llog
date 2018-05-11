package processor

import (
	"github.com/syyongx/llog/types"
	"os"
)

var ProcessId = func(record *types.Record) {
	record.Extra["Pid"] = os.Getpid()
}

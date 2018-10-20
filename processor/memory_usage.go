package processor

import (
	"github.com/syyongx/llog/types"
	"runtime"
)

// Get memory usage.
var MemoryUsage = func(record *types.Record) {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	record.Extra["MemoryUsage"] = stat.Alloc
}

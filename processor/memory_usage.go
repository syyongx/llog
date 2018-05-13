package processor

import (
	"runtime"
	"github.com/syyongx/llog/types"
)

// Get memory usage.
var MemoryUsage = func(record *types.Record) {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	record.Extra["MemoryUsage"] = stat.Alloc
}

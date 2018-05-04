package processor

import (
	"github.com/syyongx/llog/types"
)

type Processor func(record types.Record) types.Record

package types

type Processor func(record *Record, a ...interface{})

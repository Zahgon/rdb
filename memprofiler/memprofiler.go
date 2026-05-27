package memprofiler

import (
	"github.com/hdt3213/rdb/model"
)

// used to evaluate memory usage

// RedisMeta stores redis version and architecture
type RedisMeta struct {
	Version string
	Bits    int // 32/64
}

// SizeOfObject evaluates memory usage of obj
func SizeOfObject(obj model.RedisObject) int {
	_ = "STUB: not implemented"
	// todo: memory profile by redis version and architecture
	return 0
}

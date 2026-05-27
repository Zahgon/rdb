package helper

import (
	"github.com/bytedance/sonic"
)

var jsonEncoder = sonic.ConfigDefault

// ConcurrentOption sets the number of goroutines for json converter
type ConcurrentOption int

// WithConcurrent sets the number of goroutines for json converter
func WithConcurrent(c int) ConcurrentOption {
	_ = "STUB: not implemented"
	return *new(ConcurrentOption)
}

// ToJsons read rdb file and convert to json file
func ToJsons(rdbFilename string, jsonFilename string, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// open file

// create decoder

// parse rdb

// parse options

// leave one core for parser

// parser goroutine

// json marshaller goroutine

// enable SortMapKeys to ensure same result

// write goroutine

// wait writing goroutine

// finish json

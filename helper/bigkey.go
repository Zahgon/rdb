package helper

import (
	"os"
)

// FindBiggestKeys read rdb file and find the largest N keys.
// The invoker owns output, FindBiggestKeys won't close it
func FindBiggestKeys(rdbFilename string, topN int, output *os.File, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

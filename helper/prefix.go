package helper

import (
	"os"
)

// PrefixAnalyse read rdb file and find the largest N keys.
// The invoker owns output, FindBiggestKeys won't close it
func PrefixAnalyse(rdbFilename string, topN int, maxDepth int, output *os.File, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// for root(depth==1) and database root(depth==2)

// decode rdb file

// prefix tree

// get top list

// skip root and database root

// write into csv

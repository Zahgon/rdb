package helper

import (
	"github.com/hdt3213/rdb/d3flame"
	"github.com/hdt3213/rdb/model"
)

// TrimThreshold is the min count of keys to enable trim
var TrimThreshold = 1000

// FlameGraph draws flamegraph in web page to analysis memory usage pattern
func FlameGraph(rdbFilename string, port int, separators []string, options ...interface{}) (chan<- struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default port

func split(s string, separators []string) []string { _ = "STUB: not implemented"; return nil }

func addObject(root *d3flame.FlameItem, separators []string, object model.RedisObject) {
	_ = "STUB: not implemented"
	return
}

// bigNodeThreshold is the min size
var bigNodeThreshold = 1024 * 1024 // 1MB

func trimData(root *d3flame.FlameItem) {
	_ = "STUB: not implemented"
	// trim long tail
	return
}

// Aggregate leaf nodes

// child is a leaf node
// remove small leaf keys

// reserve big key

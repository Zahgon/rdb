package helper

import (
	"os"
)

type prefixStats struct {
	size     int
	keyCount int
}

// SepPrefixAnalyse reads an RDB file and aggregates memory usage by key prefix
// using a flat map (constant memory). Keys are split by the given separators up to maxDepth.
// Multiple separators are normalized to the first one before splitting.
func SepPrefixAnalyse(rdbFilename string, topN int, maxDepth int, separators []string, output *os.File, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// flat map: "db\x00prefix" -> stats

// normalize all separators to the primary one

// only emit prefixes that actually group keys —
// skip depth == len(parts) since that's the full key, not a prefix

// sort by size descending

// write CSV

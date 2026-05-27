package helper

import (
	"os"

	"github.com/hdt3213/rdb/model"
)

type hotKeyEntry struct {
	object model.RedisObject
	freq   int64
}

type hotKeyList struct {
	list     []*hotKeyEntry
	capacity int
}

func (hl *hotKeyList) add(entry *hotKeyEntry) { _ = "STUB: not implemented"; return }

func newHotKeyList(cap int) *hotKeyList { _ = "STUB: not implemented"; return nil }

// FindHotKeys read rdb file and find the hottest N keys by LFU frequency.
//
// IMPORTANT: This function only works when the RDB file was generated from a Redis instance
// configured with LFU eviction policy (maxmemory-policy allkeys-lfu or volatile-lfu).
// Under other eviction policies, the RDB file does not contain LFU frequency data,
// and the result will be empty.
//
// Keys without LFU information are skipped.
// The invoker owns output, FindHotKeys won't close it.
func FindHotKeys(rdbFilename string, topN int, output *os.File, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// no LFU info, skip

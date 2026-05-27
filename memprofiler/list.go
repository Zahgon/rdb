package memprofiler

import (
	"github.com/hdt3213/rdb/model"
)

func sizeOfListObject(obj *model.ListObject) int { _ = "STUB: not implemented"; return 0 }

func zipListIntEntryOverhead(v int) int { _ = "STUB: not implemented"; return 0 }

func zipListStrEntryOverhead(v string) int { _ = "STUB: not implemented"; return 0 }

func sizeOfZiplist(values [][]byte) int {
	_ = "STUB: not implemented"
	// See https://github.com/antirez/redis/blob/unstable/src/ziplist.c
	// <zlbytes><zltail><zllen><entry><entry><zlend>
	return 0
}

// not int

// is int

func sizeOfQuicklist(detail *model.QuicklistDetail) int { _ = "STUB: not implemented"; return 0 }

func sizeOfQuicklist2(values [][]byte, detail *model.Quicklist2Detail) int {
	_ = "STUB: not implemented"
	return 0
}

// https://github.com/CN-annotation-team/redis7.0-chinese-annotated/blob/7.0-cn-annotated/src/quicklist.h#L60

// listpack overhead: <total_bytes><size>...<end>

func sizeOfList(values [][]byte) int {
	_ = "STUB: not implemented"
	// See https://github.com/antirez/redis/blob/unstable/src/adlist.h
	// A list has 5 pointers + an unsigned long
	return 0
}

// A node has 3 pointers

// fixme: since redis 4.0, make it compatible with older version

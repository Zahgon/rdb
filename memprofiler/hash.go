package memprofiler

import "github.com/hdt3213/rdb/model"

func hashTableEntryOverhead() int {
	_ = "STUB: not implemented"
	// See  https://github.com/antirez/redis/blob/unstable/src/dict.h
	// Each dictEntry has 2 pointers + int64
	return 0
}

func hashtableOverhead(size int) int {
	_ = "STUB: not implemented"
	// See  https://github.com/antirez/redis/blob/unstable/src/dict.h
	// See the structures dict and dictht
	// 2 * (3 unsigned longs + 1 pointer) + int + long + 2 pointers
	//
	// Additionally, see **table in dictht
	// The length of the table is the next nextPower of 2
	// When the hashtable is rehashing, another instance of **table is created
	// Due to the possibility of rehashing during loading, we calculate the worse
	// case in which both tables are allocated, and so multiply
	// the size of **table by 1.5
	return 0
}

func sizeOfHashObject(obj *model.HashObject) int { _ = "STUB: not implemented"; return 0 }

func sizeOfSetObject(obj *model.SetObject) int { _ = "STUB: not implemented"; return 0 }

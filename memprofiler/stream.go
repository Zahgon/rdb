package memprofiler

import "github.com/hdt3213/rdb/model"

func sizeOfStreamObject(obj *model.StreamObject) int { _ = "STUB: not implemented"; return 0 }

// size of stream struct
// rax struct

// size of 2 new streamID and a uint64

// size of struct streamCG

// size of new field entries_read

//  streamNACK

// streamConsumer

func sizeOfStreamRaxTree(elementCount int) int {
	_ = "STUB: not implemented"
	// This is a very rough estimation. The only alternative to doing an estimation,
	// is to fully build a radix tree of similar design, and elementCount the nodes.
	// There should be at least as many nodes as there are elements in the radix tree (possibly up to 3 times)
	return 0
}

// formula for memory estimation copied from Redis's streamRadixTreeMemoryUsage

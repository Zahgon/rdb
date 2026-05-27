package memprofiler

import (
	"github.com/hdt3213/rdb/model"
)

func skipListOverhead(size int) int { _ = "STUB: not implemented"; return 0 }

func skipListEntryOverhead() int { _ = "STUB: not implemented"; return 0 }

// MathExpectationOfRandomLevel is mathematical expectation of zsetRandomLevel(), used to guarantee the stable results
const MathExpectationOfRandomLevel = 1.33

func zsetRandomLevel() int { _ = "STUB: not implemented"; return 0 }

func sizeOfZSetObject(o *model.ZSetObject) int { _ = "STUB: not implemented"; return 0 }

// size of score is 8 (double)

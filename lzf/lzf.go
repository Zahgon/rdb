package lzf

import "errors"

const (
	htabLog  uint32 = 14
	htabSize uint32 = 1 << htabLog
	maxLit          = 1 << 5
	maxOff          = 1 << 13
	maxRef          = (1 << 8) + (1 << 3)
)

var (
	errInsufficientBuffer = errors.New("insufficient buffer")
	errDataCorruption     = errors.New("data corruption")
)

// using https://github.com/zhuyie/golzf according to MIT license
// Decompress decompress lzf compressed data
func Decompress(input []byte, inLen int, outLen int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* literal run */

/* back reference */

// Can't use copy(...) here, because it has special handling when source and destination overlap.

// Compress compress data using lzf algorithm
func Compress(input []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

/* start run */

/* match found at *ref++ */

/* first a faster conservative test */

/* second the exact but rare test */

/* stop run */

/* undo run if length is zero */

/* len is now #octets - 1 */

/* start run */

/* one more literal byte we must copy */

/* stop run */
/* start run */

/* at most 3 bytes can be missing here */

/* stop run */
/* start run */

/* end run */
/* undo run if length is zero */

package core

import (
	"github.com/hdt3213/rdb/model"
)

const (
	zipStr06B = 0
	zipStr14B = 1
	zipStr32B = 2

	zipInt04B = 0x0f        // high 4 bits of Int 04 encoding
	zipInt08B = 0xfe        // 11111110
	zipInt16B = 0xc0 | 0<<4 // 11000000
	zipInt24B = 0xc0 | 3<<4 // 11110000
	zipInt32B = 0xc0 | 1<<4 // 11010000
	zipInt64B = 0xc0 | 2<<4 //11100000

	zipBigPrevLen = 0xfe
)

func (dec *Decoder) readList() ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (dec *Decoder) readQuickList() ([][]byte, *model.QuicklistDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// readQuickList2 returns
func (dec *Decoder) readQuickList2() ([][]byte, *model.Quicklist2Detail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (enc *Encoder) WriteListObject(key string, values [][]byte, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) tryWriteListZipList(key string, values [][]byte, options ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (enc *Encoder) writeQuickList(key string, values [][]byte, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) writeZipList(values []string) error { _ = "STUB: not implemented"; return nil }

// reserve 10 bytes for zip list header
// header(10bytes) + zl end(1byte)

package core

import (
	"github.com/hdt3213/rdb/model"
)

func (dec *Decoder) readSet() ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (dec *Decoder) readIntSet() (result [][]byte, detail *model.IntsetDetail, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (dec *Decoder) readListPackSet() ([][]byte, *model.ListpackDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (enc *Encoder) WriteSetObject(key string, values [][]byte, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) writeSetEncoding(key string, values [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) tryWriteIntSetEncoding(key string, values [][]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

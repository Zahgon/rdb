package core

import (
	"github.com/hdt3213/rdb/model"
)

func (dec *Decoder) readZSet(zset2 bool) ([]*model.ZSetEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dec *Decoder) readZipListZSet() ([]*model.ZSetEntry, *model.ZiplistDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (dec *Decoder) readListPackZSet() ([]*model.ZSetEntry, *model.ListpackDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (enc *Encoder) WriteZSetObject(key string, entries []*model.ZSetEntry, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) writeZSet2Encoding(key string, entries []*model.ZSetEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) tryWriteZipListZSet(key string, entries []*model.ZSetEntry) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

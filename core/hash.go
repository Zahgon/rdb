package core

import (
	"github.com/hdt3213/rdb/model"
)

/*
	if hlen <= ZIPMAP_VALUE_MAX_FREE
*/

func (dec *Decoder) readHashMap() (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readHashMapEx reads hash with field-level expiration for Redis 7.4+ (typeHashWithHfe / typeHashWithHfeRc).
// rc=true for 7.4 RC format (absolute TTL, no minExpire header), rc=false for 7.4 GA format (relative TTL with minExpire header).
func (dec *Decoder) readHashMapEx(rc bool) (map[string][]byte, map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Hash with HFEs. min TTL at start (7.4+), 7.4RC not included

// Value is absolute for 7.4RC

// 0 Indicates no TTL. This is common case so we keep it small.

// TTL is relative to minExpire (with +1 to avoid 0 that already taken)

// readHashMapExValkey reads hash with field-level expiration for Valkey 9+ (typeHash2).
// Valkey stores absolute expiration timestamps as int64 after each field-value pair.
// -1 (or negative) means no TTL, which is normalized to 0.
func (dec *Decoder) readHashMapExValkey() (map[string][]byte, map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// valkey uses -1 to indicate no TTL

func (dec *Decoder) readZipMapHash() (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//todo: scan once
// record current cursor

// recover cursor at begin position of first zip map entry

// return: len, free, error
func readZipMapEntryLen(buf []byte, cursor *int, readFree bool) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func readZipMapEntry(buf []byte, cursor *int, readFree bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip free bytes

func countZipMapEntries(buf []byte, cursor *int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// reset cursor

func (dec *Decoder) readZipListHash() (map[string][]byte, *model.ZiplistDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (dec *Decoder) readListPackHash() (map[string][]byte, *model.ListpackDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (dec *Decoder) readListPackHashEx(rc bool) (map[string][]byte, map[string]int64, *model.ListpackDetail, error) {
	_ = "STUB: not implemented"

	// This value was serialized for future use-case of streaming the object directly to FLASH (while keeping in mem its next expiration time)
	return nil, nil, nil, nil
}

func (enc *Encoder) WriteHashMapObject(key string, hash map[string][]byte, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) WriteHashMapObjectEx(key string, hash map[string][]byte, expire map[string]int64, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) writeHashEncoding(key string, hash map[string][]byte, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) writeHashEncodingEx(key string, hash map[string][]byte, expire map[string]int64, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash with HFEs. min TTL at start (7.4+), 7.4RC not included
// minExpire is the minimum non-zero expiration time across all fields,
// used as a base for relative TTL encoding. Matches Redis hashTypeGetMinExpire().

// No field has expiration, use 0

// 0 indicates no TTL

// TTL is relative to minExpire (with +1 to avoid 0 that already taken)

func (enc *Encoder) writeHash2Encoding(key string, hash map[string][]byte, expire map[string]int64, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// -1 means no TTL

func (enc *Encoder) tryWriteZipListHashMap(key string, hash map[string][]byte, options ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

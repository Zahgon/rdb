// Package core is RDB core core
package core

import (
	"bufio"
	"io"

	"github.com/hdt3213/rdb/model"
)

// Decoder is an instance of rdb parsing process
type Decoder struct {
	input     *bufio.Reader
	readCount int
	buffer    []byte

	withSpecialOpCode bool
	withSpecialTypes  map[string]ModuleTypeHandleFunc

	valkey     bool
	rdbVersion int
}

// NewDecoder creates a new RDB decoder
func NewDecoder(reader io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

// WithSpecialOpCode enables returning model.AuxObject to callback
func (dec *Decoder) WithSpecialOpCode() *Decoder { _ = "STUB: not implemented"; return nil }

// WithSpecialType enables returning redis module data structure to callback
func (dec *Decoder) WithSpecialType(moduleType string, f ModuleTypeHandleFunc) *Decoder {
	_ = "STUB: not implemented"
	return nil
}

var magicNumberRedis = []byte("REDIS")
var magicNumberValkey = []byte("VALKEY")

const (
	minVersion       = 1
	maxVersion       = 12
	minVersionValkey = 80
	maxVersionValkey = 80
)

const (
	opCodeSlotImport   = 243 /* Slot import state. */
	opCodeSlotInfo     = 244 /* Foreign slot info, safe to ignore. */
	opCodeFunction     = 245 /* function library data */
	opCodeModuleAux    = 247 /* Module auxiliary data. */
	opCodeIdle         = 248 /* LRU idle time. */
	opCodeFreq         = 249 /* LFU frequency. */
	opCodeAux          = 250 /* RDB aux field. */
	opCodeResizeDB     = 251 /* Hash table resize hint. */
	opCodeExpireTimeMs = 252 /* Expire time in milliseconds. */
	opCodeExpireTime   = 253 /* Old expire time in seconds. */
	opCodeSelectDB     = 254 /* DB number of the following keys. */
	opCodeEOF          = 255
)

const (
	typeString = iota
	typeList
	typeSet
	typeZset
	typeHash
	typeZset2 /* ZSET version 2 with doubles stored in binary. */
	typeModule
	typeModule2 // Module value parser should be registered with Decoder.WithSpecialType
	_
	typeHashZipMap
	typeListZipList
	typeSetIntSet
	typeZsetZipList
	typeHashZipList
	typeListQuickList
	typeStreamListPacks
	typeHashListPack
	typeZsetListPack
	typeListQuickList2
	typeStreamListPacks2
	typeSetListPack
	typeStreamListPacks3
	typeHashWithHfeRc         // rdb 12 (only redis 7.4 rc)
	typeHashListPackWithHfeRc // rdb 12 (only redis 7.4 rc)
	typeHashWithHfe           // since rdb 12 (redis 7.4)
	typeHashListPackWithHfe   // since rdb 12 (redis 7.4)

	typeHash2 = typeHashWithHfeRc // Hash with field-level expiration (Valkey 9+)
)

const (
	EB_EXPIRE_TIME_MAX     int64 = 0x0000FFFFFFFFFFFF
	EB_EXPIRE_TIME_INVALID int64 = EB_EXPIRE_TIME_MAX + 1
	HFE_MAX_ABS_TIME_MSEC  int64 = EB_EXPIRE_TIME_MAX >> 2
)

var encodingMap = map[int]string{
	typeString:                model.StringEncoding,
	typeList:                  model.ListEncoding,
	typeSet:                   model.SetEncoding,
	typeZset:                  model.ZSetEncoding,
	typeHash:                  model.HashEncoding,
	typeZset2:                 model.ZSet2Encoding,
	typeHashZipMap:            model.ZipMapEncoding,
	typeListZipList:           model.ZipListEncoding,
	typeSetIntSet:             model.IntSetEncoding,
	typeZsetZipList:           model.ZipListEncoding,
	typeHashZipList:           model.ZipListEncoding,
	typeListQuickList:         model.QuickListEncoding,
	typeStreamListPacks:       model.ListPackEncoding,
	typeStreamListPacks2:      model.ListPackEncoding,
	typeHashListPack:          model.ListPackEncoding,
	typeZsetListPack:          model.ListPackEncoding,
	typeListQuickList2:        model.QuickList2Encoding,
	typeSetListPack:           model.ListPackEncoding,
	typeHashWithHfeRc:         model.HashExEncoding,
	typeHashListPackWithHfeRc: model.ListPackExEncoding,
	typeHashWithHfe:           model.HashExEncoding,
	typeHashListPackWithHfe:   model.ListPackExEncoding,
	// typeHash2: model.HashExEncoding, // same 22 as typeHashWithHfeRc
}

// checkHeader checks whether input has valid RDB file header
func (dec *Decoder) checkHeader() error { _ = "STUB: not implemented"; return nil }

func (dec *Decoder) readObject(flag byte, base *model.BaseObject) (model.RedisObject, error) {
	_ = "STUB: not implemented"
	return *new(model.RedisObject), nil
}

// typeHash2 == typeHashWithHfeRc, same value 22

// Valkey 9+ Hash2: absolute timestamps after each field-value pair

// Redis 7.4: typeHashWithHfeRc (rc=true) or typeHashWithHfe (rc=false)

func (dec *Decoder) parse(cb func(object model.RedisObject) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Valkey 9+: slot info metadata, safe to skip

// safe to skip

// opcode 243: Valkey=SlotImport, Redis 8.0+=KeyMeta

// Valkey 9+: slot import state

// safe to skip

// Redis 8.0+: RDB_OPCODE_KEY_META (same opcode value 243)
// Not yet supported; return error to avoid silent data corruption

// reset expire ms

// reset lru

// reset lfu

// read crc64 at the end

// Parse parses rdb and callback
// cb returns true to continue, returns false to stop the iteration
func (dec *Decoder) Parse(cb func(object model.RedisObject) bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (dec *Decoder) GetReadCount() int { _ = "STUB: not implemented"; return 0 }

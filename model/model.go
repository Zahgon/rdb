package model

import (
	"time"
)

const (
	// StringType is redis string
	StringType = "string"
	// ListType is redis list
	ListType = "list"
	// SetType is redis set
	SetType = "set"
	// HashType is redis hash
	HashType = "hash"
	// ZSetType is redis sorted set
	ZSetType = "zset"
	// AuxType is redis metadata key-value pair
	AuxType = "aux"
	// DBSizeType is for RDB_OPCODE_RESIZEDB
	DBSizeType = "dbsize"
	// StreamType is a redis stream
	StreamType = "stream"
	// FunctionsType is redis functions
	FunctionsType = "functions"
)

const (
	// StringEncoding for string
	StringEncoding = "string"
	// ListEncoding is formed by a length encoding and some string
	ListEncoding = "list"
	// SetEncoding is formed by a length encoding and some string
	SetEncoding = "set"
	// ZSetEncoding is formed by a length encoding and some string
	ZSetEncoding = "zset"
	// HashEncoding is formed by a length encoding and some string
	HashEncoding = "hash"
	// HashExEncoding is hash with field expiration
	HashExEncoding = "hashex"
	// ZSet2Encoding is zset version2 which stores doubles in binary format
	ZSet2Encoding = "zset2"
	// ZipMapEncoding has been deprecated
	ZipMapEncoding = "zipmap"
	// ZipListEncoding  stores data in contiguous memory
	ZipListEncoding = "ziplist"
	// IntSetEncoding is a ordered list of integers
	IntSetEncoding = "intset"
	// QuickListEncoding is a list of ziplist
	QuickListEncoding = "quicklist"
	// ListPackEncoding is a new replacement for ziplist
	ListPackEncoding = "listpack"
	// ListPackExEncoding is listpack with field expiration
	ListPackExEncoding = "listpackex"
	// QuickList2Encoding is a list of listpack
	QuickList2Encoding = "quicklist2"
)

// CallbackFunc process redis object
type CallbackFunc func(object RedisObject) bool

// RedisObject is interface for a redis object
type RedisObject interface {
	// GetType returns redis type of object: string/list/set/hash/zset
	GetType() string
	// GetKey returns key of object
	GetKey() string
	// GetDBIndex returns db index of object
	GetDBIndex() int
	// GetExpiration returns expiration time, expiration of persistent object is nil
	GetExpiration() *time.Time
	// GetSize returns rdb value size in Byte
	GetSize() int
	// GetElemCount returns number of elements in list/set/hash/zset
	GetElemCount() int
	// GetEncoding returns encoding of object
	GetEncoding() string
}

// EvictionInfo is an optional interface for objects that carry LRU/LFU metadata.
// Use type assertion to check if an object implements this interface.
type EvictionInfo interface {
	// GetIdleTime returns LRU idle time of object, -1 if not available
	GetIdleTime() int64
	// GetFreq returns LFU frequency of object, -1 if not available
	GetFreq() int64
}

// BaseObject is basement of redis object
type BaseObject struct {
	DB         int         `json:"db"`                   // DB is db index of redis object
	Key        string      `json:"key"`                  // Key is key of redis object
	Expiration *time.Time  `json:"expiration,omitempty"` // Expiration is expiration time, expiration of persistent object is nil
	Size       int         `json:"size"`                 // Size is rdb value size in Byte
	Type       string      `json:"type"`                 // Type is one of string/list/set/hash/zset
	Encoding   string      `json:"encoding"`             // Encoding is the exact encoding method
	Extra      interface{} `json:"-"`                    // Extra stores more detail of encoding for memory profiler and other usages
	IdleTime   *int64      `json:"lru,omitempty"`
	Freq       *int64      `json:"lfu,omitempty"`
}

// GetKey returns key of object
func (o *BaseObject) GetKey() string {
	_ = "STUB: not implemented"

	// GetDBIndex returns db index of object
	return ""
}

func (o *BaseObject) GetDBIndex() int {
	_ = "STUB: not implemented"

	// GetEncoding returns encoding of object
	return 0
}

func (o *BaseObject) GetEncoding() string {
	_ = "STUB: not implemented"

	// GetExpiration returns expiration time, expiration of persistent object is nil
	return ""
}

func (o *BaseObject) GetExpiration() *time.Time { _ = "STUB: not implemented"; return nil }

// GetSize  returns rdb value size in Byte
func (o *BaseObject) GetSize() int {
	_ = "STUB: not implemented"

	// GetElemCount returns number of elements in list/set/hash/zset
	return 0
}

func (o *BaseObject) GetElemCount() int {
	_ = "STUB: not implemented"

	// GetIdleTime returns LRU idle time of object, -1 if not available
	return 0
}

func (o *BaseObject) GetIdleTime() int64 { _ = "STUB: not implemented"; return 0 }

// GetFreq returns LFU frequency of object, -1 if not available
func (o *BaseObject) GetFreq() int64 { _ = "STUB: not implemented"; return 0 }

// StringObject stores a string object
type StringObject struct {
	*BaseObject
	Value []byte
}

// GetType returns redis object type
func (o *StringObject) GetType() string {
	_ = "STUB: not implemented"

	// MarshalJSON marshal []byte as string
	return ""
}

func (o *StringObject) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ListObject stores a list object
type ListObject struct {
	*BaseObject
	Values [][]byte
}

// GetType returns redis object type
func (o *ListObject) GetType() string {
	_ = "STUB: not implemented"

	// GetElemCount returns number of elements in list/set/hash/zset
	return ""
}

func (o *ListObject) GetElemCount() int { _ = "STUB: not implemented"; return 0 }

// MarshalJSON marshal []byte as string
func (o *ListObject) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// HashObject stores a hash object
type HashObject struct {
	*BaseObject
	Hash             map[string][]byte
	FieldExpirations map[string]int64
}

// GetType returns redis object type
func (o *HashObject) GetType() string {
	_ = "STUB: not implemented"

	// GetElemCount returns number of elements in list/set/hash/zset
	return ""
}

func (o *HashObject) GetElemCount() int {
	_ = "STUB: not implemented"

	// MarshalJSON marshal []byte as string
	return 0
}

func (o *HashObject) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// hash/listpack with HFE

// SetObject stores a set object
type SetObject struct {
	*BaseObject
	Members [][]byte
}

// GetType returns redis object type
func (o *SetObject) GetType() string {
	_ = "STUB: not implemented"

	// GetElemCount returns number of elements in list/set/hash/zset
	return ""
}

func (o *SetObject) GetElemCount() int { _ = "STUB: not implemented"; return 0 }

// MarshalJSON marshal []byte as string
func (o *SetObject) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ZSetEntry is a key-score in sorted set
type ZSetEntry struct {
	Member string  `json:"member"`
	Score  float64 `json:"score"`
}

// ZSetObject stores a sorted set object
type ZSetObject struct {
	*BaseObject
	Entries []*ZSetEntry `json:"entries"`
}

// GetType returns redis object type
func (o *ZSetObject) GetType() string {
	_ = "STUB: not implemented"

	// GetElemCount returns number of elements in list/set/hash/zset
	return ""
}

func (o *ZSetObject) GetElemCount() int { _ = "STUB: not implemented"; return 0 }

// AuxObject stores redis metadata
type AuxObject struct {
	*BaseObject
	Value string
}

// GetType returns redis object type
func (o *AuxObject) GetType() string {
	_ = "STUB: not implemented"

	// MarshalJSON marshal []byte as string
	return ""
}

func (o *AuxObject) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// FunctionsObject stores redis functions
type FunctionsObject struct {
	*BaseObject
	FunctionsLua string `json:"functionsLua"`
}

// GetType returns redis object type
func (o *FunctionsObject) GetType() string { _ = "STUB: not implemented"; return "" }

// DBSizeObject stores db size metadata
type DBSizeObject struct {
	*BaseObject
	KeyCount uint64
	TTLCount uint64
}

// GetType returns redis object type
func (o *DBSizeObject) GetType() string {
	_ = "STUB: not implemented"

	// ModuleTypeObject stores a module type object parsed by custom handler
	return ""
}

type ModuleTypeObject struct {
	*BaseObject
	ModuleType string
	Value      interface{}
}

// GetType returns module type name
func (o *ModuleTypeObject) GetType() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON marshal []byte as string
func (o *ModuleTypeObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

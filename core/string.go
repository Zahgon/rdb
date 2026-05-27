package core

const (
	len6Bit      = 0
	len14Bit     = 1
	len32or64Bit = 2
	lenSpecial   = 3
	len32Bit     = 0x80
	len64Bit     = 0x81

	encodeInt8  = 0
	encodeInt16 = 1
	encodeInt32 = 2
	encodeLZF   = 3

	maxUint6  = 1<<6 - 1
	maxUint14 = 1<<14 - 1
	minInt24  = -1 << 23
	maxInt24  = 1<<23 - 1

	len14BitMask      byte = 0b01000000
	encodeInt8Prefix       = lenSpecial<<6 | encodeInt8
	encodeInt16Prefix      = lenSpecial<<6 | encodeInt16
	encodeInt32Prefix      = lenSpecial<<6 | encodeInt32
	encodeLZFPrefix        = lenSpecial<<6 | encodeLZF
)

// readLength parse Length Encoding
// see: https://github.com/sripathikrishnan/redis-rdb-tools/wiki/Redis-RDB-Dump-File-Format#length-encoding
func (dec *Decoder) readLength() (uint64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// get first 2 bits

func (dec *Decoder) readString() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (dec *Decoder) readInt16() (int16, error) { _ = "STUB: not implemented"; return 0, nil }

func (dec *Decoder) readInt32() (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func (dec *Decoder) readInt64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (dec *Decoder) readLiteralFloat() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (dec *Decoder) readFloat() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (dec *Decoder) readFloat32() (f float32, err error) { _ = "STUB: not implemented"; return 0, nil }

func (dec *Decoder) readLZF() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (enc *Encoder) writeLength(value uint64) error { _ = "STUB: not implemented"; return nil }

// 00 + 6 bits of data

// high 6 bit and mask(0x40)
// low 8 bit

func (enc *Encoder) writeSimpleString(s string) error { _ = "STUB: not implemented"; return nil }

func (enc *Encoder) tryWriteIntString(s string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// beyond int32 range, but within int64 range

func (enc *Encoder) writeLZFString(s string) error { _ = "STUB: not implemented"; return nil }

// write compressed length

// write uncompressed length

func (enc *Encoder) writeString(s string) error { _ = "STUB: not implemented"; return nil }

// Try LZF compression - under 20 bytes it's unable to compress even so skip it
// see rdbSaveRawString at [rdb.c](https://github.com/redis/redis/blob/unstable/src/rdb.c#L449)

// lzf may failed, while out > in

// write string without try int string. for tryWriteIntSetEncoding, writeZipList
func (enc *Encoder) writeNanString(s string) error { _ = "STUB: not implemented"; return nil }

// lzf may failed, while out > in

func (enc *Encoder) WriteStringObject(key string, value []byte, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *Encoder) writeFloat64(f float64) error { _ = "STUB: not implemented"; return nil }

// a string might be encoded as an integer, but only subset (aka uint32) of the number set
// can be encoded and then decoded back.
// e.g. the following strings can not be encoded as an integer:
// "007", "-0", "-1", "+0", "+1", "0x11"
func isEncodableUint32(s string) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

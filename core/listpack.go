package core

// readListPack returns: list of entry, list of entry size, error
func (dec *Decoder) readListPack() ([][]byte, []uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func readListPackLength(buf []byte, cursor *int) int { _ = "STUB: not implemented"; return 0 }

// list pack buf: [0, 4] -> total bytes, [4:6] -> entry count

func getBackLen(elementLen uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// readListPackEntry returns: string content, int content, entry length(encoding+content+backlen), error
func (dec *Decoder) readListPackEntry(buf []byte, cursor *int) ([]byte, int64, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

// 0xxxxxxx, uint7

// 10xxxxxx + content, string(len<=63)

// assert header == 11xxxxxx

// 110xxxxx yyyyyyyy, int13
// see https://github.com/CN-annotation-team/redis7.0-chinese-annotated/blob/fba43c524524cbdb54955a28af228b513420d78d/src/listpack.c#L586

// val is uint, must use -(8191 - val), val - 8191 will cause overflow

// 1110xxxx yyyyyyyy + content, string(len < 1<<12)

// assert header == 1111xxxx

// 11110000 aaaaaaaa bbbbbbbb cccccccc dddddddd + content, string(len < 1<<32)

// 11110001 aaaaaaaa bbbbbbbb, int16

// 11110010 aaaaaaaa bbbbbbbb cccccccc, int24

// 1111 0011 -> int32

// 11110100 8Byte -> int64

// 11111111 -> end

// readListPackEntryAsString return a string representation of entry
// It means if the entry is a integer, then format it as string
func (dec *Decoder) readListPackEntryAsString(buf []byte, cursor *int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dec *Decoder) readListPackEntryAsInt(buf []byte, cursor *int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

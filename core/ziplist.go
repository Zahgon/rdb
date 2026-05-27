package core

func (dec *Decoder) readZipList() ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (dec *Decoder) readZipListEntry(buf []byte, cursor *int) (result []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeZipListEntry(prevLen uint32, val string) []byte { _ = "STUB: not implemented"; return nil }

// encode prevLen

// try int encoding

// use int encoding

// bytes.Buffer never failed

// use string encoding

// 00 + xxxxxx

func isEncodableUint64(s string) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

package core

func readBytes(buf []byte, cursor *int, size int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readByte(buf []byte, cursor *int) (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func readZipListLength(buf []byte, cursor *int) int { _ = "STUB: not implemented"; return 0 }

// zip list buf: [0, 4] -> zlbytes, [4:8] -> zltail, [8:10] -> zllen

func (dec *Decoder) readByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (dec *Decoder) readFull(buf []byte) error { _ = "STUB: not implemented"; return nil }

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandString create a random string no longer than n
func RandString(n int) string { _ = "STUB: not implemented"; return "" }

func unsafeBytes2Str(b []byte) string { _ = "STUB: not implemented"; return "" }

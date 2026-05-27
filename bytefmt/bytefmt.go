package bytefmt

// from: github.com/cloudfoundry/bytefmt by Apache License

import (
	"errors"
)

const (
	sizeByte = 1 << (10 * iota)
	sizeKilo
	sizeMega
	sizeGiga
	sizeTera
	sizePeta
	sizeExa
)

var errInvalidByteQuantity = errors.New("byte quantity must be a positive integer with a unit of measurement like M, MB, MiB, G, GiB, or GB")

// FormatSize returns a human-readable byte string of the form 10M, 12.5K, and so forth.  The following units are available:
//
//	E: Exabyte
//	P: Petabyte
//	T: Terabyte
//	G: Gigabyte
//	M: Megabyte
//	K: Kilobyte
//	B: Byte
//
// The unit that results in the smallest number greater than or equal to 1 is always chosen.
func FormatSize(bytes uint64) string { _ = "STUB: not implemented"; return "" }

// ParseSize parses a string formatted by FormatSize as bytes. Note binary-prefixed and SI prefixed units both mean a base-2 units
// KB = K = KiB = 1024
// MB = M = MiB = 1024 * K
// GB = G = GiB = 1024 * M
// TB = T = TiB = 1024 * G
// PB = P = PiB = 1024 * T
// EB = E = EiB = 1024 * P
func ParseSize(s string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// no unit

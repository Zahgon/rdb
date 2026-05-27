// Package crc64jones implements a 64-bit cyclic redundancy check, or CRC-64,
// checksum. Specifically the Jones flavour of it which is used by Redis.
//
// Specification of this CRC64 variant follows:
// - Name: crc-64-jones
// - Width: 64 bites
// - Poly: 0xad93d23594c935a9
// - Reflected In: True
// - Xor_In: 0xffffffffffffffff
// - Reflected_Out: True
// - Xor_Out: 0x0
// - Check("123456789"): 0xe9c6d914c4b8d9ca
package crc64jones

import (
	"hash"
	"hash/crc64"
	"sync"
)

// Predefined polynomials.
const (
	// The Jones polynomial.
	Jones = 0xad93d23594c935a9
)

var table = crc64.MakeTable(reflect(Jones))

// reflect reverses the bit order of the given polynomial.
func reflect(poly uint64) uint64 { _ = "STUB: not implemented"; return 0 }

var (
	slicing8TablesBuildOnce sync.Once
	slicing8TableJones      *[8]crc64.Table
)

func buildSlicing8TablesOnce() { _ = "STUB: not implemented"; return }

func buildSlicing8Tables() { _ = "STUB: not implemented"; return }

func makeSlicingBy8Table(t *crc64.Table) *[8]crc64.Table { _ = "STUB: not implemented"; return nil }

// digest represents the partial evaluation of a checksum.
type digest struct {
	crc uint64
	tab *crc64.Table
}

// New creates a new hash.Hash64 computing the CRC-64 checksum using the
// Jones polynomial. Its Sum method will lay the value out in little-endian
// byte order.
func New() hash.Hash64 { _ = "STUB: not implemented"; return *new(hash.Hash64) }

func (d *digest) Size() int { _ = "STUB: not implemented"; return 0 }

func (d *digest) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (d *digest) Reset() { _ = "STUB: not implemented"; return }

func update(crc uint64, tab *crc64.Table, p []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// Table comparison is somewhat expensive, so avoid it for small sizes

// According to the tests between various x86 and arm CPUs, 2k is a reasonable
// threshold for now. This may change in the future.

// Update using slicing-by-8

// For reminders or small sizes

// Update returns the result of adding the bytes in p to the crc.
func Update(crc uint64, tab *crc64.Table, p []byte) uint64 { _ = "STUB: not implemented"; return 0 }

func (d *digest) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d *digest) Sum64() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *digest) Sum(in []byte) []byte {
	_ = "STUB: not implemented"

	// Compared to the core hash functions we return in little endian byte order.
	return nil
}

// Checksum returns the CRC-64 checksum of data
// using the polynomial represented by the [Table].
func Checksum(data []byte, tab *crc64.Table) uint64 { _ = "STUB: not implemented"; return 0 }

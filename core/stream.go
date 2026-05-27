package core

import (
	"github.com/hdt3213/rdb/model"
)

const (
	// StreamItemFlagNone means No special flags.
	StreamItemFlagNone = 0
	// StreamItemFlagDeleted means entry was deleted
	StreamItemFlagDeleted = 1 << 0
	// StreamItemFlagSameFields means entry has the same fields as the master entry
	StreamItemFlagSameFields = 1 << 1
)

func (dec *Decoder) readStreamListPacks(version uint) (*model.StreamObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dec *Decoder) readStreamId() (*model.StreamId, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readStreamEntries read entries
func (dec *Decoder) readStreamEntries() ([]*model.StreamEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip 4Byte total-bytes + 2Byte num-elements

// readStreamEntryContent read messages in a stream entry
func (dec *Decoder) readStreamEntryContent(buf []byte, cursor *int, firstId *model.StreamId) (*model.StreamEntry, error) {
	_ = "STUB: not implemented"
	// read count
	return nil, nil
}

// read field names of master entry

// read lp count of master entry

// ms and seq may be negative

// read lp count

func (dec *Decoder) readStreamGroups(version uint) ([]*model.StreamGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read pending list

// read consumers

// WriteStreamObject writes a stream object to RDB file
func (enc *Encoder) WriteStreamObject(key string, stream *model.StreamObject, options ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Write stream type based on version

// default to version 1

// Write stream entries

// Write stream length

// Write last ID

// Write version 2+ fields if available

// Write zero ID if FirstId is nil

// Write zero ID if MaxDeletedId is nil

// Write stream groups

// writeStreamId writes a stream ID
func (enc *Encoder) writeStreamId(id *model.StreamId) error { _ = "STUB: not implemented"; return nil }

// writeStreamEntries writes stream entries
func (enc *Encoder) writeStreamEntries(entries []*model.StreamEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// Write entry header (first message ID)

// Write entry content as listpack

// writeStreamEntryContent writes a stream entry content as listpack
func (enc *Encoder) writeStreamEntryContent(entry *model.StreamEntry) error {
	_ = "STUB: not implemented"
	// Calculate total messages (including deleted ones)
	return nil
}

// Build listpack with proper backlen values

// Add count and deleted count

// Add master field names

// Add field count for master entry (this is what the decoder reads as "end flag")

// Add messages

// Calculate flag

// Check if message uses same fields as master

// Add flag

// Add message ID (relative to first message ID)

// Add field count if not same fields

// Add fields

// Use master field names order

// Add field names and values

// Add field count for this message (this is what the decoder reads as "end flag")

// Build listpack with proper backlen values

// Write the complete listpack

type listpackEntry struct {
	intVal int64
	strVal string
}

// buildListpackWithBacklen builds a proper listpack with backlen values
func (enc *Encoder) buildListpackWithBacklen(entries []listpackEntry) []byte {
	_ = "STUB: not implemented"
	return nil
}

// First pass: encode entries and calculate sizes

// Second pass: add backlen values

// Add backlen

// Add entry

// Add header
// 6 bytes for header

// encodeBacklen encodes a backlen value
func (enc *Encoder) encodeBacklen(elementLen uint32) []byte { _ = "STUB: not implemented"; return nil }

// writeStreamGroups writes stream groups
func (enc *Encoder) writeStreamGroups(groups []*model.StreamGroup, version uint) error {
	_ = "STUB: not implemented"
	return nil
}

// Write group name

// Write last ID

// Write entries read (version 2+)

// Write pending list

// Write message ID

// Write delivery time

// Write delivery count

// Write consumers

// Write consumer name

// Write seen time

// Write active time (version 3+)

// Write consumer pending list

// Write message ID

// encodeListPackInt encodes an integer for listpack
func (enc *Encoder) encodeListPackInt(val int64) []byte { _ = "STUB: not implemented"; return nil }

// 0xxxxxxx, uint7

// 110xxxxx yyyyyyyy, int13

// 11110001 aaaaaaaa bbbbbbbb, int16

// 11110010 aaaaaaaa bbbbbbbb cccccccc, int24

// 11110011 aaaaaaaa bbbbbbbb cccccccc dddddddd, int32

// 11110100 8Byte -> int64

// encodeListPackString encodes a string for listpack
func (enc *Encoder) encodeListPackString(s string) []byte { _ = "STUB: not implemented"; return nil }

// 10xxxxxx + content, string(len<=63)

// 1110xxxx yyyyyyyy + content, string(len < 1<<12)

// 11110000 aaaaaaaa bbbbbbbb cccccccc dddddddd + content, string(len < 1<<32)

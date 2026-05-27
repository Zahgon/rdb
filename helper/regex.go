package helper

import (
	"regexp"

	"github.com/hdt3213/rdb/model"
)

type decoder interface {
	Parse(cb func(object model.RedisObject) bool) error
}

type regexDecoder struct {
	reg *regexp.Regexp
	dec decoder
}

func (d *regexDecoder) Parse(cb func(object model.RedisObject) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// regexWrapper returns
func regexWrapper(d decoder, expr string) (*regexDecoder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RegexOption enable regex filters
type RegexOption *string

// WithRegexOption creates a WithRegexOption from regex expression
func WithRegexOption(expr string) RegexOption {
	_ = "STUB: not implemented"

	// noExpiredDecoder filter all expired keys
	return *new(RegexOption)
}

type noExpiredDecoder struct {
	dec decoder
}

func (d *noExpiredDecoder) Parse(cb func(object model.RedisObject) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// NoExpiredOption tells decoder to filter all expired keys
type NoExpiredOption bool

// WithNoExpiredOption tells decoder to filter all expired keys
func WithNoExpiredOption() NoExpiredOption { _ = "STUB: not implemented"; return *new(NoExpiredOption) }

type ExpirationOption string

func WithExpirationOption(expr string) ExpirationOption {
	_ = "STUB: not implemented"
	return *new(ExpirationOption)
}

// SizeOption filters by object size
type SizeOption string

// WithSizeOption creates a SizeOption from size expression
// expression format: "<min>~<max>", supports KB/MB/GB/TB/PB/EB units and 'inf'
func WithSizeOption(expr string) SizeOption {
	_ = "STUB: not implemented"
	return *

	// expirationDecoder returns entries with expiration times and expiration within the range.
	new(SizeOption)
}

type expirationDecoder struct {
	dec             decoder
	expirationRange []int64
}

// Parse returns entries with expiration times and expiration within the range.

func (d *expirationDecoder) Parse(cb func(object model.RedisObject) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// sizeDecoder returns entries with size within the range.
type sizeDecoder struct {
	dec       decoder
	sizeRange []int
}

func (d *sizeDecoder) Parse(cb func(object model.RedisObject) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// noExpirationDecoder returns entries without expiration
type noExpirationDecoder struct {
	dec decoder
}

func (d *noExpirationDecoder) Parse(cb func(object model.RedisObject) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func parseExpireExpr(s string) ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

func parseSizeExpr(s string) ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

type GlobalMetaOption bool

func WithGlobalMeta() GlobalMetaOption { _ = "STUB: not implemented"; return *new(GlobalMetaOption) }

func wrapDecoder(dec decoder, options ...interface{}) (decoder, error) {
	_ = "STUB: not implemented"
	return *new(decoder), nil
}

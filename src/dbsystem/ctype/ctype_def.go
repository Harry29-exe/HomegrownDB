package ctype

import "HomegrownDB/common/random"

type Operations interface {
	Equal(v1, v2 []byte) bool
	Cmp(v1, v2 []byte) int
}

type Reader interface {
	// Skip skips the data belonging to this column
	// this function supports toasts and lobs
	Skip(data []byte) []byte
	// Copy copies all column data to dest
	// this function supports toasts and lobs
	Copy(dest []byte, data []byte) (copiedBytes int)

	// Value returns data in normalized format
	// (e.g. for string 4 bytes len + str, regardless received storage form)
	Value(data []byte) (value []byte)
	// ValueAndSkip gives the same result (but is more efficient)
	// as calling Value(data) and Skip(data), but it does not support
	// toasts and lobs
	ValueAndSkip(data []byte) (value, next []byte)
}

type Writer interface {
	// WriteTuple rewrites ctype from old tuple/qrow to new tuple
	// returns written bytes (support toast and lob ptrs)
	WriteTuple(dest []byte, value []byte) int
	// WriteNormalized rewrites ctype from old tuple/qrow to byte slice
	// returns written bytes (don't support toast and lob ptrs)
	WriteNormalized(dest []byte, value []byte) int
}

type Debug interface {
	// ToStr Should be called on result of Value(data) as it won't always
	// work on raw data (because of support data)
	ToStr(val []byte) string
	// Rand generate random data that normally could belong to
	// column (data is generated with support data)
	Rand(r random.Random) []byte
}

// ToastStatus The Oversized-Attribute Storage Technique status
type ToastStatus = uint8

const (
	ToastNone ToastStatus = iota
	ToastStore
	ToastLob
)

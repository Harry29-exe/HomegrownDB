package ctype

import "HomegrownDB/common/random"

type Operations interface {
	Equal(v1, v2 []byte) bool
	Cmp(v1, v2 []byte) int
}

type Reader interface {
	// Skip skips the data belonging to this column
	Skip(data []byte) []byte
	// Value of this column (without support data like length)
	Value(data []byte) (value []byte)
	// ValueAndSkip gives the same result (but is more efficient)
	// as calling Value(data) and Skip(data)
	ValueAndSkip(data []byte) (value, next []byte)
	// Copy copies all column data to dest (with support data like length)
	Copy(dest []byte, data []byte) (copiedBytes int)
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

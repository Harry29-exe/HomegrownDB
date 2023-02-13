package rawtype

import (
	"HomegrownDB/lib/random"
	"io"
)

type Type interface {
	Tag() Tag
	// Validate check if given value is fulfilling provided args
	Validate(args Args, value Value) ValidateResult
	TypeReader
	TypeWriter
	TypeOperations
	TypeDebug
}

// ValidateResult validation result that contains information about saving
// value to tuple
type ValidateResult struct {
	Status ValidateStatus // Status whether value can be assigned to column with this type
	Toast  bool           // Toast whether value should be stored in toast
	Reason error          // Reason is present only when Status is false
}

type ValidateStatus int8

const (
	ValidateOk ValidateStatus = iota
	ValidateConv
	ValidateErr
)

type TypeOperations interface {
	Equal(v1, v2 []byte) bool
	Cmp(v1, v2 []byte) int
}

type TypeReader interface {
	// Skip skips the data belonging to this column
	// this function supports toasts and lobs
	Skip(data []byte) []byte
	// Copy copies all column data to dest
	// this function supports toasts and lobs
	Copy(dest []byte, data []byte) (copiedBytes int)
	IsToastPtr(data []byte) bool
	// Value returns data in normalized format
	// (e.g. for string 4 bytes len + str, regardless received storage form)
	Value(data []byte) (value []byte)
	// ValueAndSkip gives the same result (but is more efficient)
	// as calling Value(data) and Skip(data), but it does not support
	// toasts and lobs
	ValueAndSkip(data []byte) (value, next []byte)
}

// UniWriter - universal writer combines many
type UniWriter interface {
	io.Writer
	io.ByteWriter
}

type TypeWriter interface {
	WriteValue(writer UniWriter, value Value, args Args) error
}

type TypeDebug interface {
	// ToStr Should be called on result of Value(data) as it won'ColType always
	// work on raw data (because of support data)
	ToStr(val []byte) string
	// Rand generate random data that normally could belong to
	// column (data is generated with support data)
	Rand(args Args, r random.Random) []byte
}

package hgtype

import "HomegrownDB/common/random"

type Type interface {
	Tag() TypeTag
	TypeReader
	TypeWriter
	TypeOperations
	TypeDebug
}

type Args = *typeArgs

type typeArgs struct {
	Length uint32
	VarLen bool
	UTF8   bool
}

type TypeOperations interface {
	Equal(args Args, v1, v2 []byte) bool
	Cmp(args Args, v1, v2 []byte) int
}

type TypeConverter interface {
}

type TypeReader interface {
	// Skip skips the data belonging to this column
	// this function supports toasts and lobs
	Skip(args Args, data []byte) []byte
	// Copy copies all column data to dest
	// this function supports toasts and lobs
	Copy(args Args, dest []byte, data []byte) (copiedBytes int)
	IsToastPtr(args Args, data []byte) bool
	// Value returns data in normalized format
	// (e.g. for string 4 bytes len + str, regardless received storage form)
	Value(args Args, data []byte) (value []byte)
	// ValueAndSkip gives the same result (but is more efficient)
	// as calling Value(data) and Skip(data), but it does not support
	// toasts and lobs
	ValueAndSkip(args Args, data []byte) (value, next []byte)
}

type TypeWriter interface {
	// WriteTuple rewrites hgtype from old tuple/qrow to new tuple
	// returns written bytes (support toast and lob ptrs)
	WriteTuple(args Args, dest []byte, value []byte) int
	// WriteNormalized rewrites hgtype from old tuple/qrow to byte slice
	// returns written bytes (don't support toast and lob ptrs)
	//WriteNormalized(dest []byte, value []byte) int //todo not sure if this method is needed
}

type TypeDebug interface {
	// ToStr Should be called on result of Value(data) as it won't always
	// work on raw data (because of support data)
	ToStr(args Args, val []byte) string
	// Rand generate random data that normally could belong to
	// column (data is generated with support data)
	Rand(args Args, r random.Random) []byte
}

package hgtype

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/hgtype/rawtype"
)

type (
	Args = rawtype.Args
)

type ColType interface {
	Tag() rawtype.Tag
	Type() rawtype.Type
	Args() Args
	CTOperations
	CTReader
	CTWriter
	CTDebug
}

type CTOperations interface {
	Equal(v1, v2 []byte) bool
	Cmp(v1, v2 []byte) int
}

type CTReader interface {
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

type CTWriter interface {
	WriteValue(writer rawtype.UniWriter, value rawtype.Value) error
	// WriteTuple rewrites hgtype from old tuple/qrow to new tuple
	// returns written bytes (support toast and lob ptrs)
	//WriteTuple(dest []byte, value []byte) int
	// WriteValue rewrites hgtype from old tuple/qrow to byte slice
	// returns written bytes (don'ColType support toast and lob ptrs)
	//WriteValue(dest []byte, value []byte) int //todo not sure if this method is needed
}

type CTDebug interface {
	// ToStr Should be called on result of Value(data) as it won'ColType always
	// work on raw data (because of support data)
	ToStr(val []byte) string
	// Rand generate random data that normally could belong to
	// column (data is generated with support data)
	Rand(r random.Random) []byte
}

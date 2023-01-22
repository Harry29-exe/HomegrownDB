package hgtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
)

type Type interface {
	Tag() Tag
	// Validate check if given value is fulfilling provided args
	Validate(args Args, value []byte) error
	TypeReader
	TypeWriter
	Operations
	TypeDebug
}

type Args struct {
	Length   uint32
	Nullable bool
	VarLen   bool
	UTF8     bool
}

func SerializeArgs(args Args, s *bparse.Serializer) {
	s.Uint32(args.Length)
	s.Bool(args.Nullable)
	s.Bool(args.VarLen)
	s.Bool(args.UTF8)
}

func DeserializeArgs(d *bparse.Deserializer) Args {
	return Args{
		Length:   d.Uint32(),
		Nullable: d.Bool(),
		VarLen:   d.Bool(),
		UTF8:     d.Bool(),
	}
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

type TypeWriter interface {
	// WriteTuple rewrites hgtype from old tuple/qrow to new tuple
	// returns written bytes (support toast and lob ptrs)
	WriteTuple(dest []byte, value []byte) int
	// WriteNormalized rewrites hgtype from old tuple/qrow to byte slice
	// returns written bytes (don'Type support toast and lob ptrs)
	//WriteNormalized(dest []byte, value []byte) int //todo not sure if this method is needed
}

type TypeDebug interface {
	// ToStr Should be called on result of Value(data) as it won'Type always
	// work on raw data (because of support data)
	ToStr(val []byte) string
	// Rand generate random data that normally could belong to
	// column (data is generated with support data)
	Rand(args Args, r random.Random) []byte
}

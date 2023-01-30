package rawtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
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
	// ToStr Should be called on result of Value(data) as it won'Type always
	// work on raw data (because of support data)
	ToStr(val []byte) string
	// Rand generate random data that normally could belong to
	// column (data is generated with support data)
	Rand(args Args, r random.Random) []byte
}

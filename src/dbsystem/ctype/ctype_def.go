package ctype

type Type interface {
	Reader
	Operations
	Debug
}

type Operations interface {
	Equal(v1, v2 []byte) bool
	Cmp(v1, v2 []byte) int
}

type Reader interface {
	Skip(data []byte) []byte
	Value(data []byte) (value []byte)
	ValueAndSkip(data []byte) (value, next []byte)
	Copy(data []byte, dest []byte) (copiedBytes int)
}

type Debug interface {
	ToStr(val []byte) string
}

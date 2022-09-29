package ctype

func NewStrArgs(length uint32, lengthVarying, utf8 bool) Args {
	return map[string]any{
		ArgLength:        length,
		ArgLengthVarying: lengthVarying,
		ArgUTF8:          utf8,
	}
}

// Args are additional args to create column it can be
// for example length
type Args = map[string]any

type ArgName = string

const (
	ArgLength        = "Length"
	ArgLengthVarying = "LengthVarying"
	ArgUTF8          = "UTF8"
)

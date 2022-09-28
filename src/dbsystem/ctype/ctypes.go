package ctype

type CType = uint8

const (
	Int8 CType = iota
	Float8
	Str
)

func CTypeToStr(cType CType) string {
	return ctypeNames[cType]
}

var ctypeNames map[CType]string = map[CType]string{
	Int8:   "Int8",
	Float8: "Float8",
	Str:    "Str",
}

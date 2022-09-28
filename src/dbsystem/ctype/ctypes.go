package ctype

type Type = uint8

const (
	TypeInt8 Type = iota
	Float8
	TypeStr
)

func CTypeToStr(cType Type) string {
	return ctypeNames[cType]
}

var ctypeNames map[Type]string = map[Type]string{
	TypeInt8: "TypeInt8",
	Float8:   "Float8",
	TypeStr:  "TypeStr",
}

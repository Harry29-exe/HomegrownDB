package hgtype

type Type uint8

const (
	TypeInt8 Type = iota
	TypeFloat8
	TypeStr
)

func CTypeToStr(cType Type) string {
	return ctypeNames[cType]
}

var ctypeNames = map[Type]string{
	TypeInt8:   "TypeInt8",
	TypeFloat8: "TypeFloat8",
	TypeStr:    "TypeStr",
}

type TypeArgs struct {
}

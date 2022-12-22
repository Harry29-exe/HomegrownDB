package hgtype

type TypeTag uint8

const (
	TypeInt8 TypeTag = iota
	TypeFloat8
	TypeStr
)

func (t TypeTag) ToStr() string {
	return ctypeNames[t]
}

var ctypeNames = map[TypeTag]string{
	TypeInt8:   "TypeInt8",
	TypeFloat8: "TypeFloat8",
	TypeStr:    "TypeStr",
}

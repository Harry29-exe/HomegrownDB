package hgtype

type Tag uint8

const (
	TypeInt8 Tag = iota
	TypeFloat8
	TypeStr
)

func (t Tag) ToStr() string {
	return ctypeNames[t]
}

var ctypeNames = map[Tag]string{
	TypeInt8:   "TypeInt8",
	TypeFloat8: "TypeFloat8",
	TypeStr:    "TypeStr",
}

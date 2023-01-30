package rawtype

type Tag uint8

const (
	TypeInt8 Tag = iota
	TypeFloat8
	TypeStr
)

func (t Tag) ToStr() string {
	return ctypeNames[t]
}

func (t Tag) Type() Type {
	switch t {
	case TypeInt8:
		return Int8{}
	case TypeFloat8:
		//todo implement me
		panic("Not implemented")
	case TypeStr:
		return Str{}
	default:
		panic("unknown tag: " + t.ToStr())
	}
}

var ctypeNames = map[Tag]string{
	TypeInt8:   "TypeInt8",
	TypeFloat8: "TypeFloat8",
	TypeStr:    "TypeStr",
}

package ctype

func convInputStr(input string, cType Type) ([]byte, error) {
	switch cType {
	case TypeStr:
		return []byte(input), nil
	default:
		return nil, NewInputConvErr(InputTypeStr, input, cType)
	}
}
